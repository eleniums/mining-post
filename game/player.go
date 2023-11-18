package game

type Player struct {
	Name      string  `json:"name"`
	Money     float64 `json:"money"`
	Salary    float64 `json:"salary"`
	Inventory []*Item `json:"inventory"`
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:      name,
		Money:     200000.0,
		Salary:    1000.0,
		Inventory: []*Item{},
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
