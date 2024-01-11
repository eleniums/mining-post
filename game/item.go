package game

import (
	"github.com/eleniums/mining-post/data"
)

type Item struct {
	Resource *Resource

	Quantity int64
}

// Create an item from a listing with zero quantity.
func NewItem(l *Listing) *Item {
	return &Item{
		Resource: l.Resource,
		Quantity: 0,
	}
}

// Map a database item to a game item.
func NewItemFromDB(dbItem data.Item) *Item {
	listing := stockMasterList[dbItem.Name]
	item := NewItem(listing)
	item.Quantity = dbItem.Quantity
	return item
}

// Map game item to a database item.
func (i *Item) ToDB() data.Item {
	return data.Item{
		Name:     i.Resource.Name,
		Quantity: i.Quantity,
	}
}
