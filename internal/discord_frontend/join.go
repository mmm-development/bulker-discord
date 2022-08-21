package discord_frontend

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mmm-development/bulker-discord/internal/clog"
	"github.com/mmm-development/bulker-discord/locale"
	bend "github.com/mmm-development/bulker-discord/pkg/bulker_backend"
)

var (
	BJoin_Name = "b-join"
)

func BJoin_Interaction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	code := BNew_Sessions.NewPlayer(i.GuildID, i.Member.User.ID)
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
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: locale.L.Get(i.Locale, "NewGameSession_JoinOK"),
			Flags:   uint64(discordgo.MessageFlagsEphemeral),
		},
	})

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
