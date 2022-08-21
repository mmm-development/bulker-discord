package bend

type GameSessionReturnCode int

const (
	OK                      GameSessionReturnCode = iota
	IS_A_HOST                                     = iota
	PLAYER_EXISTS                                 = iota
	PLAYER_NOT_EXISTS                             = iota
	GAME_SESSION_EXISTS                           = iota
	GAME_SESSION_NOT_EXISTS                       = iota
)

func (gsrc GameSessionReturnCode) LocaleKey() string {
	switch gsrc {
	case IS_A_HOST:
		return "NewGameSession_IsAHost"
	case PLAYER_EXISTS:
		return "NewGameSession_PlayerExists"
	case PLAYER_NOT_EXISTS:
		return "NewGameSession_PlayerNotExists"
	case GAME_SESSION_EXISTS:
		return "NewGameSession_GameSessionExists"
	case GAME_SESSION_NOT_EXISTS:
		return "NewGameSession_GameSessionNotExists"
	}
	return ""
}
