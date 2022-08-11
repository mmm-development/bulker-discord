package appcmd

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/mmm-development/bulker-discord/bend"
	"github.com/mmm-development/bulker-discord/locale"
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

	BNew_Sessions = make(bend.GameSessionMap)
)

func BNew_Message(guildID string, i *discordgo.InteractionCreate) *discordgo.MessageSend {
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
				Title: locale.L.Get(i.Locale, "BNew_Caption"),
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  locale.L.Get(i.Locale, "BNew_HostCaption"),
						Value: fmt.Sprintf("<@%s>", gsdata.Host),
					},
					{
						Name:  locale.L.Get(i.Locale, "BNew_PlayersCaption"),
						Value: playersList,
					},
				},
			},
		},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    locale.L.Get(i.Locale, "BNew_JoinButton"),
						Style:    discordgo.SuccessButton,
						Disabled: false,
						CustomID: "b-join",
					},
					discordgo.Button{
						Label:    locale.L.Get(i.Locale, "BNew_QuitButton"),
						Style:    discordgo.DangerButton,
						Disabled: false,
						CustomID: "b-leave",
					},
				},
			},
		},
	}
}

func BNew_Interaction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	m, err := s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Color: 0xFFE299,
				Title: locale.L.Get(i.Locale, "BNew_Caption"),
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  locale.L.Get(i.Locale, "BNew_HostCaption"),
						Value: fmt.Sprintf("<@%s>", i.Member.User.ID),
					},
					{
						Name:  locale.L.Get(i.Locale, "BNew_PlayersCaption"),
						Value: "-",
					},
				},
			},
		},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    locale.L.Get(i.Locale, "BNew_JoinButton"),
						Style:    discordgo.SuccessButton,
						Disabled: false,
						CustomID: "b-join",
					},
					discordgo.Button{
						Label:    locale.L.Get(i.Locale, "BNew_QuitButton"),
						Style:    discordgo.DangerButton,
						Disabled: false,
						CustomID: "b-leave",
					},
				},
			},
		},
	})
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "failed to create game session :(",
				Flags:   uint64(discordgo.MessageFlagsEphemeral),
			},
		})
		fmt.Printf("[ERROR] Creating game session:")
		fmt.Println(err)
	} else {
		if BNew_Sessions.NewGameSession(i.GuildID, i.Member.User.ID) != bend.OK {

		}

		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
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
								Label:    locale.L.Get(i.Locale, "BNew_HostStartButton"),
								Style:    discordgo.SuccessButton,
								Disabled: false,
								CustomID: "b-start",
							},
							discordgo.Button{
								Emoji: discordgo.ComponentEmoji{
									Name: "❌",
								},
								Label:    locale.L.Get(i.Locale, "BNew_HostCancelButton"),
								Style:    discordgo.DangerButton,
								Disabled: false,
								CustomID: "b-abort",
							},
						},
					},
				},
			},
		})

		if err != nil {
			fmt.Printf("[ERROR] Creating host menu:")
			fmt.Println(err)
		}
	}
}
