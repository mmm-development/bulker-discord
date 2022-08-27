package bend

type GameGatherMap map[string]*GameGather

func (ggm *GameGatherMap) NewGameGather(guildID string, hostID string) GameGatherReturnCode {
	if _, ok := (*ggm)[guildID]; ok {
		return GAME_SESSION_EXISTS
	}
	(*ggm)[guildID] = NewGameGather(hostID)
	return OK
}

func (ggm *GameGatherMap) CleanGameGather(guildID string) (*GameGatherData, GameGatherReturnCode) {
	if _, ok := (*ggm)[guildID]; !ok {
		return nil, GAME_SESSION_NOT_EXISTS
	}
	gsData := (*ggm)[guildID].GetPlayers()
	delete(*ggm, guildID)
	return gsData, OK
}

func (ggm *GameGatherMap) NewPlayer(guildID string, id string) GameGatherReturnCode {
	if _, ok := (*ggm)[guildID]; !ok {
		return GAME_SESSION_NOT_EXISTS
	}
	return (*ggm)[guildID].NewPlayer(id)
}

func (ggm *GameGatherMap) DeletePlayer(guildID string, id string) GameGatherReturnCode {
	if _, ok := (*ggm)[guildID]; !ok {
		return GAME_SESSION_NOT_EXISTS
	}
	return (*ggm)[guildID].DeletePlayer(id)
}

func (ggm GameGatherMap) GetPlayers(guildID string) (*GameGatherData, GameGatherReturnCode) {
	if _, ok := ggm[guildID]; !ok {
		return nil, GAME_SESSION_NOT_EXISTS
	}
	return ggm[guildID].GetPlayers(), OK
}
