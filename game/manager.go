package game

import (
	"fmt"
	"log/slog"
	"sync"
	"time"
)

const updateInterval = 10 * time.Second

type Manager struct {
	market     Market // TODO: maybe move this to be a sync.Map of just the listings
	players    *sync.Map
	ticker     *time.Ticker
	marketLock sync.RWMutex
	playerLock *sync.Map
}

func NewManager() *Manager {
	manager := &Manager{
		players:    &sync.Map{},
		playerLock: &sync.Map{},
	}

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
		manager.playerLock.Store(name, &sync.RWMutex{})
	}

	return manager
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
	// TODO: going to have to lock each market listing... otherwise will have to lock the entire market to make a sale
	m.marketLock.RLock()
	defer m.marketLock.RUnlock()

	return DeepCopy(m.market.Stock)
}

func (m *Manager) GetPlayer(name string) (Player, error) {
	lock, ok := MapLoad[string, *sync.RWMutex](m.playerLock, name)
	if !ok {
		return Player{}, fmt.Errorf("error locking for player: %s", name)
	}
	defer lock.RLock()

	player, ok := MapLoad[string, *Player](m.players, name)
	if !ok {
		return Player{}, fmt.Errorf("player does not exist with name: %s", name)
	}

	return *player, nil
}
