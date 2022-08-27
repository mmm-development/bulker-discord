package discord_frontend

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mmm-development/bulker-discord/internal/clog"
	"github.com/mmm-development/bulker-discord/locale"
	bend "github.com/mmm-development/bulker-discord/pkg/bulker_backend"
)

var (
	BAbort_Name = "b-abort"
)

func BAbort_Interaction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	BNew_SessionStartSignal[i.GuildID] <- struct{}{}
	_, code := BNew_Sessions.CleanGameGather(i.GuildID)

	if code != bend.OK {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: locale.L.Get(i.Locale, code.LocaleKey()),
				Flags:   uint64(discordgo.MessageFlagsEphemeral),
			},
		})
		clog.L.Info("Aborting game session:\n%s", locale.L.Get(locale.DefLocale, code.LocaleKey()))
		return
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: locale.L.Get(i.Locale, "NewGameGather_AbortOK"),
			Flags:   uint64(discordgo.MessageFlagsEphemeral),
		},
	})
}
