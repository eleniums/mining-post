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
