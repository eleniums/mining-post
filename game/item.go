package game

type Item struct {
	Resource

	Quantity int64 `json:"quantity"`

	// If set, this function will be called every world update. Useful for
	// generating materials or something extra from equipment.
	update func(player *Player, item *Item)
}
