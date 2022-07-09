package locale

import "github.com/bwmarrin/discordgo"

func init() {
	L[discordgo.Russian] = make(map[string]string)
	var ru map[string]string = L[discordgo.Russian]

	ru["BPing_Name"] = "б-пинг"
	ru["BPing_Description"] = "Тестовая команда для проверки работоспособности бота"
	ru["BPing_Response"] = "Я жив!"
}
