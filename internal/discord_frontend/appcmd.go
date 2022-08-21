package discord_frontend

import "github.com/bwmarrin/discordgo"

var (
	Commands = []*discordgo.ApplicationCommand{
		&BPing_AppCmd,
		&BNew_AppCmd,
	}
	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		BPing_Name: BPing_Interaction,
		BNew_Name:  BNew_Interaction,
	}
	ComponentHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		BAbort_Name: BAbort_Interaction,
		BJoin_Name:  BJoin_Interaction,
		BLeave_Name: BLeave_Interaction,
		BStart_Name: BStart_Interaction,
	}
)
