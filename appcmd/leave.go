package appcmd

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mmm-development/bulker-discord/bend"
	"github.com/mmm-development/bulker-discord/clog"
	"github.com/mmm-development/bulker-discord/locale"
)

var (
	BLeave_Name = "b-leave"
)

func BLeave_Interaction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	code := BNew_Sessions.DeletePlayer(i.GuildID, i.Member.User.ID)
	if code != bend.OK {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: locale.L.Get(i.Locale, code.LocaleKey()),
				Flags:   uint64(discordgo.MessageFlagsEphemeral),
			},
		})
		return
	}

	msg := BNew_Message(i.GuildID, i.Locale)
	_, err := s.ChannelMessageEditComplex(&discordgo.MessageEdit{
		ID:         BNew_SessionInitMsg[i.GuildID],
		Channel:    i.ChannelID,
		Content:    &msg.Content,
		Components: msg.Components,
		Embeds:     msg.Embeds,
	})
	if err != nil {
		clog.L.Error("Editing game session message:\n%v", err)
	}
}
