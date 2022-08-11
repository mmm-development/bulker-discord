package bend

type GameSession struct {
	Host   string
	Joined map[string]struct{}
}

type GameSessionData struct {
	Host   string
	Joined []string
}

func NewGameSession(hostID string) *GameSession {
	return &GameSession{
		Host:   hostID,
		Joined: make(map[string]struct{}),
	}
}

func (gs *GameSession) NewPlayer(id string) GameSessionReturnCode {
	if id == gs.Host {
		return IS_A_HOST
	}
	if _, ok := gs.Joined[id]; ok {
		return PLAYER_EXISTS
	}
	gs.Joined[id] = struct{}{}
	return OK
}

func (gs *GameSession) DeletePlayer(id string) GameSessionReturnCode {
	if _, ok := gs.Joined[id]; !ok {
		return PLAYER_NOT_EXISTS
	}
	delete(gs.Joined, id)
	return OK
}

func (gs GameSession) GetPlayers() *GameSessionData {
	playersList := make([]string, len(gs.Joined))
	i := 0
	for player := range gs.Joined {
		playersList[i] = player
	}

	return &GameSessionData{
		Host:   gs.Host,
		Joined: playersList,
	}
}
