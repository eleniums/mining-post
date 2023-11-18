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

// Add quantity of item to player's inventory.
func (p *Player) AddItem(item *Item) {
	// if item already exists, just add to quantity
	found, ok := Find(p.Inventory, item, func(a *Item, b *Item) bool {
		return a.Name == b.Name
	})
	if ok {
		found.Quantity += item.Quantity
		return
	}

	// if item doesn't exist, add it
	p.Inventory = append(p.Inventory, item)
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
