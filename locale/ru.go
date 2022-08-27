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

	ru["NewGameGather_StartOK"] = "Игра успешно начата!"
	ru["NewGameGather_AbortOK"] = "Игра успешно отменена."
	ru["NewGameGather_JoinOK"] = "Вы записаны в игру!"
	ru["NewGameGather_LeaveOK"] = "Вы вышли из игры."
	ru["NewGameGather_IsAHost"] = "Вы не можете присоединиться к игре как игрок. Вы ведь ведущий!"
	ru["NewGameGather_PlayerExists"] = "Вы уже в игре!"
	ru["NewGameGather_PlayerNotExists"] = "Вы уже вне игры!"
	ru["NewGameGather_GameGatherExists"] = "На сервере уже создана игровая сессия."
	ru["NewGameGather_GameGatherNotExists"] = "Игровая сессия либо началась, либо была отменена."
}
