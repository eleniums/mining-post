package game

import (
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/eleniums/mining-post/data"
)

const updateInterval = 10 * time.Second

type Manager struct {
	NextUpdate time.Time

	startTime time.Time

	market  map[string]*Listing
	players map[string]*Player

	worldLock *sync.RWMutex
	ticker    *time.Ticker

	db Storage
}

func NewManager(db Storage) (*Manager, error) {
	manager := &Manager{
		startTime: time.Now(),
		market:    map[string]*Listing{},
		players:   map[string]*Player{},
		worldLock: &sync.RWMutex{},
		db:        db,
	}

	// create market listings
	manager.market = resourceMap

	// load player data from database
	dbPlayers, err := manager.db.LoadPlayers()
	if err != nil {
		return nil, err
	}
	for _, dbPlayer := range dbPlayers {
		player := NewPlayerFromDB(dbPlayer)
		manager.players[dbPlayer.Name] = player
	}

	// add in test users if they don't already exist
	testPlayers := loadTestPlayers()
	for _, player := range testPlayers {
		if _, ok := manager.players[player.Name]; !ok {
			manager.players[player.Name] = player
		}
	}

	// randomize prices for all listings for the initial loop
	manager.adjustMarketPrices()

	return manager, nil
}

// Start game engine with regular updates.
func (m *Manager) Start() {
	slog.Info("Initialize the game and start updates", "loopInterval", updateInterval)
	m.ticker = time.NewTicker(updateInterval)
	m.NextUpdate = time.Now().Add(updateInterval).UTC()
	go func() {
		for range m.ticker.C {
			m.NextUpdate = time.Now().Add(updateInterval).UTC()
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
	slog.Info("Game update starting...", "serverStarted", m.startTime.UTC(), "serverUptime", time.Since(m.startTime))

	startTime := time.Now()

	// stop the world while updating
	m.worldLock.Lock()
	defer m.worldLock.Unlock()

	// randomize all market prices and quantities for the next round
	m.adjustMarketPrices()

	// run updates for each player
	playerSaveData := make([]data.Player, len(m.players))
	playerSaveIndex := 0
	for _, player := range m.players {
		player.lock.Lock()
		defer player.lock.Unlock()

		// give player salary
		player.Money += player.Salary

		// run updates on any resources as needed
		for _, item := range player.Inventory {
			if item.Resource.update != nil {
				item.Resource.update(player, item)
			}
		}

		// calculate player net worth
		player.NetWorth = player.Money
		for _, item := range player.Inventory {
			if listing, ok := m.market[item.Resource.Name]; ok {
				player.NetWorth += float64(item.Quantity) * listing.SellPrice
			}
		}

		// check for player promotion
		if ranks[player.Rank].eligibleForPromotion(player) {
			player.Rank++
			player.Title = ranks[player.Rank].Name
			player.Salary = ranks[player.Rank].Salary
			slog.Info("Player was promoted to a new rank", "username", player.Name, "title", player.Title, "rank", player.Rank)
		}

		// prepare player data to be saved
		playerSaveData[playerSaveIndex] = player.ToDB()
		playerSaveIndex++
	}

	// save all player data
	err := m.db.SavePlayers(playerSaveData)
	if err != nil {
		slog.Error("Failed to save player data", ErrAttr(err))
	}

	slog.Info("Game update finished", "playerCount", len(m.players), "elapsed", time.Since(startTime))
}

// Used internally to adjust all market prices.
func (m *Manager) adjustMarketPrices() {
	for _, listing := range m.market {
		listing.adjustMarketPrice()
	}
}

func (m *Manager) GetMarketStock(filters ...ListingFilter) []*Listing {
	m.worldLock.RLock()
	defer m.worldLock.RUnlock()

	listings := MapValues(m.market)

	for _, filter := range filters {
		switch filter.Property {
		case LISTING_FILTER_NAME:
			listings = Filter(listings, func(val *Listing) bool {
				return val.Resource.Name == filter.Value
			})
		case LISTING_FILTER_TYPE:
			listings = Filter(listings, func(val *Listing) bool {
				return string(val.Resource.Type) == filter.Value
			})
		}
	}

	return CopySlice(listings)
}

func (m *Manager) GetPlayer(name string) (*Player, error) {
	player, ok := m.players[name]
	if !ok {
		return nil, fmt.Errorf("player does not exist with name: %s", name)
	}

	player.lock.RLock()
	defer player.lock.RUnlock()

	return Copy(player), nil
}

func (m *Manager) BuyOrder(playerName string, resourceName string, quantity int64) (float64, error) {
	if quantity <= 0 {
		return 0, errors.New("quantity must be greater than 0")
	}

	player, ok := m.players[playerName]
	if !ok {
		return 0, fmt.Errorf("player does not exist with name: %s", playerName)
	}

	m.worldLock.RLock()
	defer m.worldLock.RUnlock()

	player.lock.Lock()
	defer player.lock.Unlock()

	listing, ok := m.market[resourceName]
	if !ok {
		return 0, fmt.Errorf("resource not found for purchase: %s", resourceName)
	}

	// determine cost of resource at requested quantity
	cost := listing.BuyPrice * float64(quantity)

	// determine if player can afford to purchase the requested quantity
	if player.Money < cost {
		return 0, fmt.Errorf("insufficient funds to purchase %d of resource: %s, cost: %.2f", quantity, resourceName, cost)
	}

	// TODO: map prereqs to responses, so you can see them in market listing (not inventory though)
	// check if player meets prerequisites to purchase resource
	for _, v := range listing.Resource.Prerequisites {
		item := player.GetResource(v.Name)
		if item == nil || item.Quantity < v.Quantity*quantity {
			return 0, fmt.Errorf("player does not meet prerequisites to purchase resource: %s", resourceName)
		}
	}

	// remove prerequisite resources from player's inventory
	for _, v := range listing.Resource.Prerequisites {
		player.AddResource(player.GetResource(v.Name).Resource, -v.Quantity*quantity)
	}

	// buy resource for player
	player.Money -= cost
	err := player.AddResource(listing.Resource, quantity)
	if err != nil {
		return 0, err
	}

	return cost, nil
}

func (m *Manager) SellOrder(playerName string, resourceName string, quantity int64) (float64, error) {
	if quantity <= 0 {
		return 0, errors.New("quantity must be greater than 0")
	}

	player, ok := m.players[playerName]
	if !ok {
		return 0, fmt.Errorf("player does not exist with name: %s", playerName)
	}

	m.worldLock.RLock()
	defer m.worldLock.RUnlock()

	player.lock.Lock()
	defer player.lock.Unlock()

	listing, ok := m.market[resourceName]
	if !ok {
		return 0, fmt.Errorf("resource not found for sale: %s", resourceName)
	}

	// determine profit of resource at requested quantity
	profit := listing.SellPrice * float64(quantity)

	// determine if player has enough quantity to sell
	item := player.GetResource(resourceName)
	if item == nil {
		return 0, fmt.Errorf("player does not have resource in inventory: %s", resourceName)
	}
	if item.Quantity < quantity {
		return 0, fmt.Errorf("insufficient quantity to sell %d of resource: %s", quantity, resourceName)
	}

	// sell resource for player
	player.Money += profit

	// remove quantity from player's inventory
	err := player.AddResource(item.Resource, -quantity)
	if err != nil {
		return 0, err
	}

	return profit, nil
}
