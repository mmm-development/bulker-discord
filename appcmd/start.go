package appcmd

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/mmm-development/bulker-discord/bend"
	"github.com/mmm-development/bulker-discord/clog"
	"github.com/mmm-development/bulker-discord/locale"
)

var (
	BStart_Name = "b-start"
)

func BStart_Interaction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	playersSet, code := BNew_Sessions.CleanGameSession(i.GuildID)
	if code != bend.OK {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: locale.L.Get(i.Locale, code.LocaleKey()),
				Flags:   uint64(discordgo.MessageFlagsEphemeral),
			},
		})
		clog.L.Info("Starting game session:\n%s", locale.L.Get(locale.DefLocale, code.LocaleKey()))
		return
	}

	if _, ok := BNew_SessionInitMsg[i.GuildID]; ok {
		err := s.ChannelMessageDelete(i.ChannelID, BNew_SessionInitMsg[i.GuildID])
		if err != nil {
			clog.L.Error("Deleting game session message:\n%v", err)
		}
		delete(BNew_SessionInitMsg, i.GuildID)
	}

	for _, playerID := range playersSet.Joined {
		userSession, err := s.UserChannelCreate(playerID)
		if err != nil {
			clog.L.Error("Responding to user <@%s>:\n%v", playerID, err)
			continue
		}
		s.ChannelMessageSend(userSession.ID, ":eye:")
	}

	hostSession, err := s.UserChannelCreate(playersSet.Host)
	if err != nil {
		clog.L.Error("Responding to user <@%s>:\n%v", playersSet.Host, err)
		return
	}
	s.ChannelMessageSend(hostSession.ID, fmt.Sprintf("Собрали для тебя этих игроков:\n%sКак тебе?", strings.Join(playersSet.Joined, "\n")))
}
