package bend

type GameGather struct {
	Host   string
	Joined map[string]struct{}
}

type GameGatherData struct {
	Host   string
	Joined []string
}

func NewGameGather(hostID string) *GameGather {
	return &GameGather{
		Host:   hostID,
		Joined: make(map[string]struct{}),
	}
}

func (gg *GameGather) NewPlayer(id string) GameGatherReturnCode {
	if id == gg.Host {
		return IS_A_HOST
	}
	if _, ok := gg.Joined[id]; ok {
		return PLAYER_EXISTS
	}
	gg.Joined[id] = struct{}{}
	return OK
}

func (gg *GameGather) DeletePlayer(id string) GameGatherReturnCode {
	if _, ok := gg.Joined[id]; !ok {
		return PLAYER_NOT_EXISTS
	}
	delete(gg.Joined, id)
	return OK
}

func (gg GameGather) GetPlayers() *GameGatherData {
	playersList := make([]string, len(gg.Joined))
	i := 0
	for player := range gg.Joined {
		playersList[i] = player
	}

	return &GameGatherData{
		Host:   gg.Host,
		Joined: playersList,
	}
}
