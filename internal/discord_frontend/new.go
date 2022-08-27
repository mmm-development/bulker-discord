package discord_frontend

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/mmm-development/bulker-discord/internal/clog"
	"github.com/mmm-development/bulker-discord/locale"
	bend "github.com/mmm-development/bulker-discord/pkg/bulker_backend"
)

var (
	BNew_Name   = "b-new"
	BNew_AppCmd = discordgo.ApplicationCommand{
		Name:                     BNew_Name,
		Description:              "Create new game session",
		Type:                     discordgo.ChatApplicationCommand,
		NameLocalizations:        locale.L.LocaleMap("BNew_Name"),
		DescriptionLocalizations: locale.L.LocaleMap("BNew_Description"),
	}

	BNew_Sessions           = make(bend.GameGatherMap)
	BNew_SessionInitMsg     = make(map[string]string)
	BNew_SessionStartSignal = make(map[string]chan struct{})
)

func BNew_Message(guildID string, userLocale discordgo.Locale) *discordgo.MessageSend {
	gsdata, statusCode := BNew_Sessions.GetPlayers(guildID)
	if statusCode != bend.OK {
		return nil
	}

	playersList := ""
	for _, name := range gsdata.Joined {
		playersList += fmt.Sprintf("<@%s>\n", name)
	}
	if len(playersList) == 0 {
		playersList = "-"
	}

	return &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Color: 0xFFE299,
				Title: locale.L.Get(userLocale, "BNew_Caption"),
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  locale.L.Get(userLocale, "BNew_HostCaption"),
						Value: fmt.Sprintf("<@%s>", gsdata.Host),
					},
					{
						Name:  locale.L.Get(userLocale, "BNew_PlayersCaption"),
						Value: playersList,
					},
				},
			},
		},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    locale.L.Get(userLocale, "BNew_JoinButton"),
						Style:    discordgo.SuccessButton,
						Disabled: false,
						CustomID: BJoin_Name,
					},
					discordgo.Button{
						Label:    locale.L.Get(userLocale, "BNew_QuitButton"),
						Style:    discordgo.DangerButton,
						Disabled: false,
						CustomID: BLeave_Name,
					},
				},
			},
		},
	}
}

func BNew_ModeratorMessage(userLocale discordgo.Locale) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: uint64(discordgo.MessageFlagsEphemeral),
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Emoji: discordgo.ComponentEmoji{
								Name: "✔️",
							},
							Label:    locale.L.Get(userLocale, "BNew_HostStartButton"),
							Style:    discordgo.SuccessButton,
							Disabled: false,
							CustomID: BStart_Name,
						},
						discordgo.Button{
							Emoji: discordgo.ComponentEmoji{
								Name: "❌",
							},
							Label:    locale.L.Get(userLocale, "BNew_HostCancelButton"),
							Style:    discordgo.DangerButton,
							Disabled: false,
							CustomID: BAbort_Name,
						},
					},
				},
			},
		},
	}
}

func BNew_Interaction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var err error
	var statusCode bend.GameGatherReturnCode
	var st *discordgo.Message
	var msg *discordgo.MessageSend

	statusCode = BNew_Sessions.NewGameGather(i.GuildID, i.Member.User.ID)
	if statusCode != bend.OK {
		goto ON_CREATE_ERROR_INTERACTION
	}

	msg = BNew_Message(i.GuildID, i.Locale)
	if msg == nil {
		err = errors.New("failed to create game session message")
		goto ON_INIT_ERROR_INTERACTION
	}

	err = s.InteractionRespond(i.Interaction, BNew_ModeratorMessage(i.Locale))
	if err != nil {
		goto ON_INIT_ERROR_INTERACTION
	}

	st, err = s.ChannelMessageSendComplex(i.ChannelID, msg)
	if err != nil {
		goto ON_INIT_ERROR_INTERACTION
	}
	BNew_SessionInitMsg[i.GuildID] = st.ID

	BNew_SessionStartSignal[i.GuildID] = make(chan struct{})
	<-BNew_SessionStartSignal[i.GuildID]
	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{Content: "-", Components: []discordgo.MessageComponent{}})
	if err != nil {
		clog.L.Error("Cleaning game session control panel:\n%v", err)
	}
	close(BNew_SessionStartSignal[i.GuildID])
	delete(BNew_SessionStartSignal, i.GuildID)
	delete(BNew_SessionInitMsg, i.GuildID)
	err = s.ChannelMessageDelete(i.ChannelID, st.ID)
	if err != nil {
		clog.L.Error("Deleting game session message:\n%v", err)
	}
	return

ON_CREATE_ERROR_INTERACTION:
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: locale.L.Get(i.Locale, statusCode.LocaleKey()),
			Flags:   uint64(discordgo.MessageFlagsEphemeral),
		},
	})
	clog.L.Info("Creating game session:\n%s", locale.L.Get(locale.DefLocale, statusCode.LocaleKey()))
	return

ON_INIT_ERROR_INTERACTION:
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "failed to create game session :(",
			Flags:   uint64(discordgo.MessageFlagsEphemeral),
		},
	})
	clog.L.Error("Creating game session:\n%v", err)
	BNew_Sessions.CleanGameGather(i.GuildID)
	if st != nil {
		err = s.ChannelMessageDelete(i.ChannelID, st.ID)
		if err != nil {
			clog.L.Error("Deleting game session message:\n%v", err)
		}
	}
}
