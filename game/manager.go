package game

import (
	"fmt"
	"log/slog"
	"sync"
	"time"
)

const updateInterval = 10 * time.Second

type Manager struct {
	market  Market
	players *sync.Map
	ticker  *time.Ticker
}

func NewManager() *Manager {
	// create market
	market := Market{
		Stock: GetInitialStock(),
	}

	// randomize all listings for the initial loop
	market.Randomize()

	// add players
	players := &sync.Map{}
	players.Store("snelson", Player{
		Name:      "snelson",
		Money:     100.00,
		Inventory: []Item{},
	})

	return &Manager{
		market:  market,
		players: players,
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

// Update happens on a regular time interval. This is the main game loop.
func (m *Manager) update() {
	startTime := time.Now()

	// as a last step, randomize all market prices and quantities for the next round
	m.market.Randomize()

	slog.Info("Game update finished", "elapsed", time.Since(startTime))
}

func (m *Manager) GetMarketStock() Market {
	return m.market
}

func (m *Manager) GetPlayer(name string) (Player, error) {
	player, ok := m.players.Load(name)
	if !ok {
		return Player{}, fmt.Errorf("player does not exist with name: %s", name)
	}
	return player.(Player), nil
}
