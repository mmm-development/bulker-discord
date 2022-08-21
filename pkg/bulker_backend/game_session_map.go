package bend

type GameSessionMap map[string]*GameSession

func (gsm *GameSessionMap) NewGameSession(guildID string, hostID string) GameSessionReturnCode {
	if _, ok := (*gsm)[guildID]; ok {
		return GAME_SESSION_EXISTS
	}
	(*gsm)[guildID] = NewGameSession(hostID)
	return OK
}

func (gsm *GameSessionMap) CleanGameSession(guildID string) (*GameSessionData, GameSessionReturnCode) {
	if _, ok := (*gsm)[guildID]; !ok {
		return nil, GAME_SESSION_NOT_EXISTS
	}
	gsData := (*gsm)[guildID].GetPlayers()
	delete(*gsm, guildID)
	return gsData, OK
}

func (gsm *GameSessionMap) NewPlayer(guildID string, id string) GameSessionReturnCode {
	if _, ok := (*gsm)[guildID]; !ok {
		return GAME_SESSION_NOT_EXISTS
	}
	return (*gsm)[guildID].NewPlayer(id)
}

func (gsm *GameSessionMap) DeletePlayer(guildID string, id string) GameSessionReturnCode {
	if _, ok := (*gsm)[guildID]; !ok {
		return GAME_SESSION_NOT_EXISTS
	}
	return (*gsm)[guildID].DeletePlayer(id)
}

func (gsm GameSessionMap) GetPlayers(guildID string) (*GameSessionData, GameSessionReturnCode) {
	if _, ok := gsm[guildID]; !ok {
		return nil, GAME_SESSION_NOT_EXISTS
	}
	return gsm[guildID].GetPlayers(), OK
}
