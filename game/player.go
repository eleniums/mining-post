package game

import (
	"fmt"
)

type Player struct {
	Name      string  `json:"name"`
	Money     float64 `json:"money"`
	Inventory []Item  `json:"inventory"`
}

type Item struct {
	Resource

	Quantity int64 `json:"quantity"`
}

func (m *Manager) GetPlayer(name string) (Player, error) {
	player, ok := m.players.Load(name)
	if !ok {
		return Player{}, fmt.Errorf("player does not exist with name: %s", name)
	}
	return player.(Player), nil
}
