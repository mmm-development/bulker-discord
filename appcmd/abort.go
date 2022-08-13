package appcmd

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mmm-development/bulker-discord/bend"
	"github.com/mmm-development/bulker-discord/clog"
	"github.com/mmm-development/bulker-discord/locale"
)

var (
	BAbort_Name = "b-abort"
)

func BAbort_Interaction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	_, code := BNew_Sessions.CleanGameSession(i.GuildID)
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

	if _, ok := BNew_SessionInitMsg[i.GuildID]; ok {
		err := s.ChannelMessageDelete(i.ChannelID, BNew_SessionInitMsg[i.GuildID])
		if err != nil {
			clog.L.Error("Deleting game session message:\n%v", err)
		}
		delete(BNew_SessionInitMsg, i.GuildID)
	}
}
