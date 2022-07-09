package locale

import "github.com/bwmarrin/discordgo"

func init() {
	L[discordgo.EnglishUS] = make(map[string]string)
	var en_us map[string]string = L[discordgo.EnglishUS]

	en_us["BPing_Name"] = "b-ping"
	en_us["BPing_Description"] = "The test command used for health check"
	en_us["BPing_Response"] = "I'm alive!"
}
