package locale

import "github.com/bwmarrin/discordgo"

func init() {
	L[discordgo.EnglishUS] = make(map[string]string)
	var en_us map[string]string = L[discordgo.EnglishUS]

	en_us["BPing_Name"] = "b-ping"
	en_us["BPing_Description"] = "The test command used for health check"
	en_us["BPing_Response"] = "I'm alive!"

	en_us["BNew_Name"] = "b-create"
	en_us["BNew_Description"] = "Create a new game session"
	en_us["BNew_Caption"] = "New Bunker Session"
	en_us["BNew_HostCaption"] = "Host"
	en_us["BNew_PlayersCaption"] = "Players' List"
	en_us["BNew_JoinButton"] = "I'm in!"
	en_us["BNew_QuitButton"] = "I'm out."
	en_us["BNew_HostStartButton"] = "Start the Game"
	en_us["BNew_HostCancelButton"] = "Cancel"

	en_us["NewGameSession_IsAHost"] = "You can't join the game as a player. You're the host!"
	en_us["NewGameSession_PlayerExists"] = "You're already in the game!"
	en_us["NewGameSession_PlayerNotExists"] = "You're already out of the game!"
	en_us["NewGameSession_GameSessionExists"] = "There is already a game session hosted on this server."
	en_us["NewGameSession_GameSessionNotExists"] = "That game session either has begun or had been cancelled."
}
