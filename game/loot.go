package game

// Loot table for calculating production of mines.
type LootTable []LootEntry

// Single entry in a loot table.
type LootEntry struct {
	Name      string // Name of resource to provide.
	Weight    int    // Weight (or chance) for this entry in the loot table.
	CountLow  int64  // Lowest number of resource to provide if chosen.
	CountHigh int64  // Highest number of resource to provide if chosen.
}

// Roll and return the name and quantity of loot.
func (lt LootTable) CalculateLoot() (string, int64) {
	// get total weight
	weight := 0
	for _, v := range lt {
		weight += v.Weight
	}

	// roll a random number
	roll := randInt(1, weight)

	// get associated loot entry
	prevTotal := 0
	total := 0
	for _, v := range lt {
		prevTotal = total
		total += v.Weight
		if roll > prevTotal && roll <= total {
			return v.Name, randInt64(v.CountLow, v.CountHigh)
		}
	}

	// this should not happen, no loot returned
	return "", 0
}
