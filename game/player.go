package game

type Player struct {
	Name      string  `json:"name"`
	Money     float64 `json:"money"`
	Inventory []Item  `json:"inventory"`
}

type Item struct {
	Resource

	Quantity int64 `json:"quantity"`
}

// Load players into memory.
func loadPlayers() map[string]*Player {
	return map[string]*Player{
		"snelson": {
			Name:      "snelson",
			Money:     100.00,
			Inventory: []Item{},
		},
		"tstark": {
			Name:      "tstark",
			Money:     100.00,
			Inventory: []Item{},
		},
		"hhughes": {
			Name:      "hhughes",
			Money:     100.00,
			Inventory: []Item{},
		},
		"dhayter": {
			Name:      "dhayter",
			Money:     100.00,
			Inventory: []Item{},
		},
	}
}
