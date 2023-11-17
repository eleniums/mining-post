package game

import (
	"sync"

	"github.com/eleniums/mining-post/models"
)

type Manager struct {
	market  models.Market
	players sync.Map
}

func NewManager() *Manager {
	return &Manager{
		market:  models.Market{},
		players: sync.Map{},
	}
}

func (m *Manager) Update() {

}
