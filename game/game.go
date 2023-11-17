package game

import (
	"sync"

	"github.com/eleniums/mining-post/models"
)

type GameManager struct {
	market  models.Market
	players sync.Map
}

func NewGameManager() *GameManager {
	return &GameManager{
		market:  models.Market{},
		players: sync.Map{},
	}
}
