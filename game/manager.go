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
	players sync.Map
	ticker  *time.Ticker

	marketLock       sync.RWMutex
	singlePlayerLock sync.Map
}

func NewManager() *Manager {
	var manager Manager

	// create market
	manager.market = Market{
		Stock: getInitialStock(),
	}

	// randomize all listings for the initial loop
	manager.market.Randomize()

	// load players
	players := loadPlayers()
	for name, player := range players {
		manager.players.Store(name, player)
		manager.singlePlayerLock.Store(name, sync.RWMutex{})
	}

	return &manager
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

	// stop the market while updating
	m.marketLock.Lock()
	defer m.marketLock.Unlock()

	// as a last step, randomize all market prices and quantities for the next round
	m.market.Randomize()

	slog.Info("Game update finished", "elapsed", time.Since(startTime))
}

func (m *Manager) GetMarketStock() []Listing {
	m.marketLock.RLock()
	defer m.marketLock.RUnlock()

	return DeepCopy(m.market.Stock)
}

func (m *Manager) GetPlayer(name string) (Player, error) {
	p, ok := m.players.Load(name)
	if !ok {
		return Player{}, fmt.Errorf("player does not exist with name: %s", name)
	}
	player := p.(*Player)

	return *player, nil
}
