package game

import (
	"fmt"
	"log/slog"
	"sync"
	"time"
)

const updateInterval = 10 * time.Second

type Manager struct {
	worldLock  sync.RWMutex
	market     *sync.Map
	players    *sync.Map
	playerLock *sync.Map
	ticker     *time.Ticker
}

func NewManager() *Manager {
	manager := &Manager{
		market:     &sync.Map{},
		players:    &sync.Map{},
		playerLock: &sync.Map{},
	}

	// create market listing
	listings := getInitialStock()
	for _, listing := range listings {
		manager.market.Store(listing.Name, listing)
	}

	// load players
	players := loadPlayers()
	for _, player := range players {
		manager.players.Store(player.Name, player)
		manager.playerLock.Store(player.Name, &sync.RWMutex{})
	}

	// randomize prices for all listings for the initial loop
	manager.adjustMarketPrices()

	return manager
}

// Start game engine with regular updates.
func (m *Manager) Start() {
	slog.Info("Initialize the game and start updates", "loop-interval", updateInterval)
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

	// stop the world while updating
	m.worldLock.Lock()
	defer m.worldLock.Unlock()

	// randomize all market prices and quantities for the next round
	m.adjustMarketPrices()

	// run updates for each player
	m.players.Range(func(key, val any) bool {
		player := val.(*Player)

		playerLock, ok := MapLoad[string, *sync.RWMutex](m.playerLock, player.Name)
		if !ok {
			slog.Error("error finding lock for player", "username", player.Name)
			return true
		}
		playerLock.RLock()
		defer playerLock.RUnlock()

		// give player salary
		player.Money += player.Salary

		// run updates on any items as needed
		for _, item := range player.Inventory {
			if item.update != nil {
				item.update(player, item)
			}
		}
		return true
	})

	slog.Info("Game update finished", "elapsed", time.Since(startTime))
}

func (m *Manager) adjustMarketPrices() {
	m.market.Range(func(key, val any) bool {
		listing := val.(*Listing)
		listing.adjustMarketPrice()
		return true
	})
}

func (m *Manager) GetMarketStock() []*Listing {
	m.worldLock.RLock()
	defer m.worldLock.RUnlock()

	return MapFlatten[string, *Listing](m.market)
}

func (m *Manager) GetPlayer(name string) (*Player, error) {
	playerLock, ok := MapLoad[string, *sync.RWMutex](m.playerLock, name)
	if !ok {
		return nil, fmt.Errorf("error finding lock for player: %s", name)
	}
	playerLock.RLock()
	defer playerLock.RUnlock()

	player, ok := MapLoad[string, *Player](m.players, name)
	if !ok {
		return nil, fmt.Errorf("player does not exist with name: %s", name)
	}

	return player, nil
}
