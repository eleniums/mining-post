package game

import (
	"github.com/eleniums/mining-post/data"
)

// Inventory item which is a quantity of some resource.
type Item struct {
	Resource *Resource
	Quantity int64
}

// Create an inventory item from a market listing with given quantity.
func NewItem(l *Listing, quantity int64) *Item {
	return &Item{
		Resource: l.Resource,
		Quantity: quantity,
	}
}

// Map a database item to a game item.
func NewItemFromDB(dbItem data.Item) *Item {
	listing := stockMasterList[dbItem.Name]
	item := NewItem(listing, dbItem.Quantity)
	return item
}

// Map game item to a database item.
func (i *Item) ToDB() data.Item {
	return data.Item{
		Name:     i.Resource.Name,
		Quantity: i.Quantity,
	}
}
