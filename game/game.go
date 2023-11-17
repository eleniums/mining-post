package game

import (
	"log/slog"
	"sync"
	"time"

	"github.com/eleniums/mining-post/models"
)

const updateInterval = 10 * time.Second

type Manager struct {
	market  models.Market
	players sync.Map
	ticker  *time.Ticker
}

func NewManager() *Manager {
	return &Manager{
		market:  models.Market{},
		players: sync.Map{},
	}
}

// Start game engine with regular updates.
func (m *Manager) Start() {
	slog.Info("Initialize the game and start updates")
	m.ticker = time.NewTicker(updateInterval)
	go func() {
		for range m.ticker.C {
			m.update()
		}
	}()
}

// Halt updates and stop the game engine.
func (m *Manager) Stop() error {
	slog.Info("Stop the game and halt updates")
	m.ticker.Stop()
	return nil
}

// Update happens on a regular time interval.
func (m *Manager) update() {
	slog.Info("Game update started")

	// TODO: implement updates

	slog.Info("Game update finished")
}
