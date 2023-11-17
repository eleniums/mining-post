package game

import (
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
		Stock: []Listing{
			{
				Resource: Resource{
					Name:        "Granite",
					Description: "Rough hewn igneous rock.",
				},
				Quantity:  1000,
				BuyPrice:  10.00,
				SellPrice: 8.00,
			},
		},
	}

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

// Update happens on a regular time interval.
func (m *Manager) update() {
	startTime := time.Now()

	// TODO: implement updates

	slog.Info("Game update finished", "elapsed", time.Since(startTime))
}
