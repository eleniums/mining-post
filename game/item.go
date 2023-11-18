package game

type Item struct {
	Resource

	Quantity int64 `json:"quantity"`

	// If set, this function will be called every world update. Useful for
	// generating materials or something extra from equipment.
	update func(player *Player, item *Item)
}

// Create an item from a listing with zero quantity.
func NewItem(l *Listing) *Item {
	return &Item{
		Resource: l.Resource,
		Quantity: 0,
		update:   l.update,
	}
}
