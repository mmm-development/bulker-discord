package appcmd

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
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

	BNew_Sessions = make(map[string]string)
)

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
		BNew_Sessions[i.GuildID] = m.ID

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
