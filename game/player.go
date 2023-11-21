package game

type Player struct {
	Name      string  `json:"name"`
	Title     string  `json:"title"`
	Rank      int     `json:"-"`
	Money     float64 `json:"money"`
	Salary    float64 `json:"salary"`
	Inventory []*Item `json:"inventory"`
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:      name,
		Title:     ranks[0].Name,
		Rank:      0,
		Money:     10000.0,
		Salary:    1000.0,
		Inventory: []*Item{},
	}
}

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

// Load players into memory.
func loadPlayers() []*Player {
	return []*Player{
		NewPlayer("snelson"),
		NewPlayer("tstark"),
		NewPlayer("hhughes"),
		NewPlayer("dhayter"),
	}
}
