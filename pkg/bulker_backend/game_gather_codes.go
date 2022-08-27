package bend

type GameGatherReturnCode int

const (
	OK                      GameGatherReturnCode = iota
	IS_A_HOST                                    = iota
	PLAYER_EXISTS                                = iota
	PLAYER_NOT_EXISTS                            = iota
	GAME_SESSION_EXISTS                          = iota
	GAME_SESSION_NOT_EXISTS                      = iota
)

func (gsrc GameGatherReturnCode) LocaleKey() string {
	switch gsrc {
	case IS_A_HOST:
		return "NewGameGather_IsAHost"
	case PLAYER_EXISTS:
		return "NewGameGather_PlayerExists"
	case PLAYER_NOT_EXISTS:
		return "NewGameGather_PlayerNotExists"
	case GAME_SESSION_EXISTS:
		return "NewGameGather_GameGatherExists"
	case GAME_SESSION_NOT_EXISTS:
		return "NewGameGather_GameGatherNotExists"
	}
	return ""
}
