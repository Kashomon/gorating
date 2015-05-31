package gorating

// Construct a map from a player id to all the games they played.
func PlayerMap(games []Game) map[string][]Game {
	m := make(map[string][]Game)
	for _, g := range games {
		for _, p := range []Player{g.PlayerOne(), g.PlayerTwo()} {
			if _, ok := m[p.UnqiueId()]; !ok {
				m[p.UnqiueId()] = make([]Game, 0, 5)
			}
			arr := m[p.UnqiueId()]
			m[p.UnqiueId()] = append(arr, g)
		}
	}
	return m
}