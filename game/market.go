package game

type Market struct {
	Stock []Listing `json:"stock"`
}

// Randomize quantities and prices of all listings.
func (m *Market) Randomize() {
	for i := range m.Stock {
		m.Stock[i].Randomize()
	}
}
