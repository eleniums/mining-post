package game

// Loot table for calculating production of mines.
type LootTable struct {
	Loot []LootEntry // Entries in the loot table.
}

// Single entry in a loot table.
type LootEntry struct {
	Name      string // Name of resource to provide.
	Weight    int    // Weight (or chance) for this entry in the loot table.
	CountLow  int    // Lowest number of resource to provide if chosen.
	CountHigh int    // Highest number of resource to provide if chosen.
}
