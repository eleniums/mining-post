package game

import (
	"fmt"
	"log/slog"
	"sync"
	"time"
)

const updateInterval = 10 * time.Second

type Manager struct {
	NextUpdate time.Time

	market  map[string]*Listing
	players map[string]*Player

	worldLock *sync.RWMutex
	ticker    *time.Ticker
}

func NewManager() *Manager {
	manager := &Manager{
		market:    map[string]*Listing{},
		players:   map[string]*Player{},
		worldLock: &sync.RWMutex{},
	}

	// create market listing
	listings := getInitialStock()
	for _, listing := range listings {
		manager.market[listing.Name] = listing
	}

	// load players
	players := loadPlayers()
	for _, player := range players {
		manager.players[player.Name] = player
	}

	// randomize prices for all listings for the initial loop
	manager.adjustMarketPrices()

	return manager
}

// Start game engine with regular updates.
func (m *Manager) Start() {
	slog.Info("Initialize the game and start updates", "loop-interval", updateInterval)
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
	startTime := time.Now()

	// stop the world while updating
	m.worldLock.Lock()
	defer m.worldLock.Unlock()

	// randomize all market prices and quantities for the next round
	m.adjustMarketPrices()

	// run updates for each player
	for _, player := range m.players {
		player.lock.Lock()
		defer player.lock.Unlock()

		// give player salary
		player.Money += player.Salary

		// run updates on any items as needed
		for _, item := range player.Inventory {
			if item.update != nil {
				item.update(player, item)
			}
		}

		// calculate player net worth
		player.NetWorth = player.Money
		for _, item := range player.Inventory {
			if listing, ok := m.market[item.Name]; ok {
				player.NetWorth += float64(item.Quantity) * listing.SellPrice
			}
		}

		// check for player promotion
		if ranks[player.Rank].eligibleForPromotion(player) {
			player.Rank++
			player.Title = ranks[player.Rank].Name
			slog.Info("player was promoted to a new rank", "username", player.Name, "title", player.Title, "rank", player.Rank)
		}
	}

	slog.Info("Game update finished", "elapsed", time.Since(startTime))
}

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
				return val.Name == filter.Value
			})
		case LISTING_FILTER_TYPE:
			listings = Filter(listings, func(val *Listing) bool {
				return string(val.Type) == filter.Value
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

func (m *Manager) BuyOrder(playerName string, itemName string, quantity int64) (float64, error) {
	player, ok := m.players[playerName]
	if !ok {
		return 0, fmt.Errorf("player does not exist with name: %s", playerName)
	}

	m.worldLock.RLock()
	defer m.worldLock.RUnlock()

	player.lock.Lock()
	defer player.lock.Unlock()

	listing, ok := m.market[itemName]
	if !ok {
		return 0, fmt.Errorf("item not found for purchase: %s", itemName)
	}

	// determine cost of item at requested quantity
	cost := listing.BuyPrice * float64(quantity)

	// determine if player can afford to purchase the requested quantity
	if player.Money < cost {
		return 0, fmt.Errorf("insufficient funds to purchase %d of item: %s, cost: %.2f", quantity, itemName, cost)
	}

	// check if player meets prerequisites to purchase item (and adjust player inventory as needed)
	if listing.prebuy != nil && !listing.prebuy(player) {
		return 0, fmt.Errorf("player does not meet prerequisites to purchase item: %s", itemName)
	}

	// buy item for player
	player.Money -= cost

	if item := player.GetItem(listing.Name); item != nil {
		// if item already exists, just add to quantity
		item.Quantity += quantity
	} else {
		// if item doesn't exist, add item to player's inventory
		item := NewItem(listing)
		item.Quantity = quantity
		player.AddItem(item)
	}

	return cost, nil
}

func (m *Manager) SellOrder(playerName string, itemName string, quantity int64) (float64, error) {
	player, ok := m.players[playerName]
	if !ok {
		return 0, fmt.Errorf("player does not exist with name: %s", playerName)
	}

	m.worldLock.RLock()
	defer m.worldLock.RUnlock()

	player.lock.Lock()
	defer player.lock.Unlock()

	listing, ok := m.market[itemName]
	if !ok {
		return 0, fmt.Errorf("item not found for sale: %s", itemName)
	}

	// determine profit of item at requested quantity
	profit := listing.SellPrice * float64(quantity)

	// determine if player has enough quantity to sell
	item := player.GetItem(itemName)
	if item == nil {
		return 0, fmt.Errorf("player does not have item in inventory: %s", itemName)
	}
	if item.Quantity < quantity {
		return 0, fmt.Errorf("insufficient quantity to sell %d of item: %s", quantity, itemName)
	}

	// check if player meets prerequisites to sell item (and adjust player inventory as needed)
	if listing.presell != nil && !listing.presell(player) {
		return 0, fmt.Errorf("player does not meet prerequisites to sell item: %s", itemName)
	}

	// sell item for player
	player.Money += profit

	// remove quantity from player's inventory
	item.Quantity -= quantity
	if item.Quantity <= 0 {
		player.RemoveItem(itemName)
	}

	return profit, nil
}
