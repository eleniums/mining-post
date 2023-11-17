package game

import (
	"fmt"

	"github.com/eleniums/mining-post/models"
)

func (m *Manager) GetPlayer(name string) (models.Player, error) {
	player, ok := m.players.Load(name)
	if !ok {
		return models.Player{}, fmt.Errorf("player does not exist with name: %s", name)
	}
	return player.(models.Player), nil
}
