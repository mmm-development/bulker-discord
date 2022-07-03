package appcmd

import "github.com/bwmarrin/discordgo"

var (
	BPing_Name   = "b-ping"
	BPing_AppCmd = discordgo.ApplicationCommand{
		Name:        BPing_Name,
		Description: "The test command used for health check",
		Type:        discordgo.ChatApplicationCommand,
		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.Russian: "б-пинг",
		},
		DescriptionLocalizations: &map[discordgo.Locale]string{
			discordgo.Russian: "Тестовая команда для проверки работоспособности бота",
		},
	}
)

func BPing_Interaction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	PingResponseLocalization := map[discordgo.Locale]string{
		discordgo.Russian:   "Я жив!",
		discordgo.EnglishUS: "I'm alive!",
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: PingResponseLocalization[i.Locale],
			Flags:   uint64(discordgo.MessageFlagsEphemeral),
		},
	})
}
