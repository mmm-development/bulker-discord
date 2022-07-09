package appcmd

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mmm-development/bulker-discord/locale"
)

var (
	BPing_Name   = "b-ping"
	BPing_AppCmd = discordgo.ApplicationCommand{
		Name:                     BPing_Name,
		Description:              "The test command used for health check",
		Type:                     discordgo.ChatApplicationCommand,
		NameLocalizations:        locale.L.LocaleMap("BPing_Name"),
		DescriptionLocalizations: locale.L.LocaleMap("BPing_Description"),
	}
)

func BPing_Interaction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: locale.L.Get(i.Locale, "BPing_Response"),
			Flags:   uint64(discordgo.MessageFlagsEphemeral),
		},
	})
}
