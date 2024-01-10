package game

import (
	"sync"

	"github.com/eleniums/mining-post/data"
)

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

// TODO: Maybe this should be how items are added or removed? Quantity can be positive or negative. Would couple it to the masterlist...
// func (p *Player) AddItem(itemName string, quantity int64) error {
// 	return nil
// }

// Add item to player's inventory.
func (p *Player) AddItem(item *Item) {
	p.Inventory = append(p.Inventory, item)
}

// Get item from player's inventory.
func (p *Player) GetItem(itemName string) *Item {
	for _, item := range p.Inventory {
		if item.Name == itemName {
			return item
		}
	}
	return nil
}

// Remove item from player's inventory.
func (p *Player) RemoveItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			return
		}
	}
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

	// TODO: add some random other players for perf testing
	// for i := 0; i < 1_000; i++ {
	// 	players = append(players, NewPlayer(strconv.Itoa(i)))
	// }

	return players
}
