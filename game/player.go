package game

import (
	"errors"
	"sync"

	"github.com/eleniums/mining-post/data"
)

// Represents a player in the game.
type Player struct {
	Name      string
	Title     string
	Rank      int
	NetWorth  float64
	Money     float64
	Salary    float64
	Inventory []*Item

	lock *sync.RWMutex
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:      name,
		Title:     ranks[0].Name,
		Rank:      0,
		NetWorth:  50.0,
		Money:     50.0,
		Salary:    10.0,
		Inventory: []*Item{},
		lock:      &sync.RWMutex{},
	}
}

// Map a database player to a game player.
func NewPlayerFromDB(dbPlayer data.Player) *Player {
	player := NewPlayer(dbPlayer.Name)
	player.Title = dbPlayer.Title
	player.Rank = dbPlayer.Rank
	player.NetWorth = dbPlayer.NetWorth
	player.Money = dbPlayer.Money
	player.Salary = dbPlayer.Salary

	player.Inventory = make([]*Item, len(dbPlayer.Inventory))
	for i, dbItem := range dbPlayer.Inventory {
		player.Inventory[i] = NewItemFromDB(dbItem)
	}

	return player
}

// Map game player to a database player.
func (p *Player) ToDB() data.Player {
	dbPlayer := data.Player{
		Name:     p.Name,
		Title:    p.Title,
		Rank:     p.Rank,
		NetWorth: p.NetWorth,
		Money:    p.Money,
		Salary:   p.Salary,
	}

	dbPlayer.Inventory = make([]data.Item, len(p.Inventory))
	for i, item := range p.Inventory {
		dbPlayer.Inventory[i] = item.ToDB()
	}

	return dbPlayer
}

// Add or remove resource quantity from player's inventory.
func (p *Player) AddResource(resource *Resource, quantity int64) error {
	for i, item := range p.Inventory {
		if item.Resource.Name == resource.Name {
			if item.Quantity+quantity < 0 {
				return errors.New("quantity cannot be negative for existing item")
			}

			// if item already exists, just add to quantity
			item.Quantity += quantity

			// if quantity is 0 or less, remove it from player's inventory
			if item.Quantity <= 0 {
				p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			}

			return nil
		}
	}

	if quantity < 1 {
		return errors.New("quantity cannot be less than 1 for new item")
	}

	// if item doesn't exist, add item to player's inventory
	item := NewItem(resource, quantity)
	p.Inventory = append(p.Inventory, item)

	return nil
}

// Get item from player's inventory.
func (p *Player) GetResource(name string) *Item {
	for _, item := range p.Inventory {
		if item.Resource.Name == name {
			return item
		}
	}
	return nil
}

// Returns some fake players to test with.
func loadTestPlayers() []*Player {
	// TODO: these are for testing purposes. Need to remove later.
	players := []*Player{
		NewPlayer("bbanner"),
		NewPlayer("tstark"),
		NewPlayer("hhughes"),
		NewPlayer("dhayter"),
	}

	// for i := 0; i < 1_000; i++ {
	// 	players = append(players, NewPlayer(strconv.Itoa(i)))
	// }

	return players
}
