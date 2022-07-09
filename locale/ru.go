package locale

import "github.com/bwmarrin/discordgo"

func init() {
	L[discordgo.Russian] = make(map[string]string)
	var ru map[string]string = L[discordgo.Russian]

	ru["BPing_Name"] = "б-пинг"
	ru["BPing_Description"] = "Тестовая команда для проверки работоспособности бота"
	ru["BPing_Response"] = "Я жив!"

	ru["BNew_Name"] = "б-создать"
	ru["BNew_Description"] = "Создать новую игровую сессию"
	ru["BNew_Caption"] = "Новая сессия Бункера"
	ru["BNew_HostCaption"] = "Ведущий"
	ru["BNew_PlayersCaption"] = "Список игроков"
	ru["BNew_JoinButton"] = "Я в деле!"
	ru["BNew_QuitButton"] = "Я пас."
	ru["BNew_HostStartButton"] = "Начать игру"
	ru["BNew_HostCancelButton"] = "Отменить"
}
