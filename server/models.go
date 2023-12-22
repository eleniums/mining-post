package server

import (
	"github.com/eleniums/mining-post/game"
)

type Player struct {
	Name      string  `json:"name"`
	Title     string  `json:"title"`
	NetWorth  float64 `json:"netWorth"`
	Money     float64 `json:"money"`
	Salary    float64 `json:"salary"`
	Inventory []Item  `json:"inventory"`
}

func NewPlayer(src *game.Player) Player {
	inventory := make([]Item, len(src.Inventory))
	for i, v := range src.Inventory {
		inventory[i] = NewItem(v)
	}

	return Player{
		Name:      src.Name,
		Title:     src.Title,
		NetWorth:  src.NetWorth,
		Money:     src.Money,
		Salary:    src.Salary,
		Inventory: inventory,
	}
}

type Resource struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

func NewResource(src game.Resource) Resource {
	return Resource{
		Name:        src.Name,
		Description: src.Description,
		Type:        string(src.Type),
	}
}

type Item struct {
	Resource

	Quantity int64 `json:"quantity"`
}

func NewItem(src *game.Item) Item {
	return Item{
		Resource: NewResource(src.Resource),
		Quantity: src.Quantity,
	}
}

type Listing struct {
	Resource

	BuyPrice  float64 `json:"buy_price"`
	SellPrice float64 `json:"sell_price"`
}

func NewListing(src *game.Listing) Listing {
	return Listing{
		Resource:  NewResource(src.Resource),
		BuyPrice:  src.BuyPrice,
		SellPrice: src.SellPrice,
	}
}
