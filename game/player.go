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

// Load players into memory.
func loadPlayers() []*Player {
	return []*Player{
		NewPlayer("snelson"),
		NewPlayer("tstark"),
		NewPlayer("hhughes"),
		NewPlayer("dhayter"),
	}
}
