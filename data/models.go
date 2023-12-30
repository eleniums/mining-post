package data

type Player struct {
	Name      string
	Title     string
	Rank      int
	NetWorth  float64
	Money     float64
	Salary    float64
	Inventory []Item
}

type Item struct {
	Name        string
	Description string
	Type        string
	Quantity    int64
}
