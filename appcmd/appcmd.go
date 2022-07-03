package appcmd

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
)
