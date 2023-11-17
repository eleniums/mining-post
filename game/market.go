package game

import (
	"sync"
)

type Market struct {
	Stock []Listing `json:"stock"`

	lock sync.RWMutex
}

// Randomize quantities and prices of all listings.
func (m *Market) Randomize() {
	// TODO: might want to move this when doing other updates
	// stop the world while updating
	m.lock.Lock()
	defer m.lock.Unlock()

	for i := range m.Stock {
		m.Stock[i].Randomize()
	}
}
