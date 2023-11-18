package game

type Player struct {
	Name      string  `json:"name"`
	Money     float64 `json:"money"`
	Inventory []*Item `json:"inventory"`
}

type Item struct {
	Resource

	Quantity int64 `json:"quantity"`

	// If set, this function will be called every world update. Useful for
	// generating materials or something extra from equipment.
	update func(player *Player, item *Item)
}

// Load players into memory.
func loadPlayers() []*Player {
	return []*Player{
		{
			Name:      "snelson",
			Money:     100.00,
			Inventory: []*Item{},
		},
		{
			Name:      "tstark",
			Money:     100.00,
			Inventory: []*Item{},
		},
		{
			Name:      "hhughes",
			Money:     100.00,
			Inventory: []*Item{},
		},
		{
			Name:      "dhayter",
			Money:     100.00,
			Inventory: []*Item{},
		},
	}
}
