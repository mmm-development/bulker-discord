package appcmd

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	BNew_Name   = "b-new"
	BNew_AppCmd = discordgo.ApplicationCommand{
		Name:        BNew_Name,
		Description: "Create new game session",
		Type:        discordgo.ChatApplicationCommand,
		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.Russian: "б-создать",
		},
		DescriptionLocalizations: &map[discordgo.Locale]string{
			discordgo.Russian: "Создать новую игровую сессию",
		},
	}

	BNew_Sessions = make(map[string]string)
)

func BNew_Interaction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	CaptionLocalization := map[discordgo.Locale]string{
		discordgo.Russian:   "Сбор в игру",
		discordgo.EnglishUS: "Game Gathering",
	}
	HostLocalization := map[discordgo.Locale]string{
		discordgo.Russian:   "Ведущий",
		discordgo.EnglishUS: "Host",
	}
	PlayersLocalization := map[discordgo.Locale]string{
		discordgo.Russian:   "Список игроков",
		discordgo.EnglishUS: "Player List",
	}
	JoinLocalization := map[discordgo.Locale]string{
		discordgo.Russian:   "Я в деле!",
		discordgo.EnglishUS: "I'm in!",
	}
	LeaveLocalization := map[discordgo.Locale]string{
		discordgo.Russian:   "Я пас.",
		discordgo.EnglishUS: "I'm out.",
	}
	StartLocalization := map[discordgo.Locale]string{
		discordgo.Russian:   "Начать игру",
		discordgo.EnglishUS: "Start the game",
	}
	AbortLocalization := map[discordgo.Locale]string{
		discordgo.Russian:   "Отменить игру",
		discordgo.EnglishUS: "Cancel the game",
	}

	m, err := s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Color: 0xFFE299,
				Title: CaptionLocalization[i.Locale],
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  HostLocalization[i.Locale],
						Value: fmt.Sprintf("<@%s>", i.Member.User.ID),
					},
					{
						Name:  PlayersLocalization[i.Locale],
						Value: "-",
					},
				},
			},
		},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    JoinLocalization[i.Locale],
						Style:    discordgo.SuccessButton,
						Disabled: false,
						CustomID: "b-join",
					},
					discordgo.Button{
						Label:    LeaveLocalization[i.Locale],
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
								Label:    StartLocalization[i.Locale],
								Style:    discordgo.SuccessButton,
								Disabled: false,
								CustomID: "b-start",
							},
							discordgo.Button{
								Emoji: discordgo.ComponentEmoji{
									Name: "❌",
								},
								Label:    AbortLocalization[i.Locale],
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
