package server

import (
	"fmt"

	"github.com/eleniums/mining-post/game"
)

type Player struct {
	Name      string `json:"name"`
	Title     string `json:"title"`
	NetWorth  string `json:"netWorth"`
	Money     string `json:"money"`
	Salary    string `json:"salary"`
	Inventory []Item `json:"inventory"`
}

func NewPlayer(src *game.Player) Player {
	inventory := make([]Item, len(src.Inventory))
	for i, v := range src.Inventory {
		inventory[i] = NewItem(v)
	}

	return Player{
		Name:      src.Name,
		Title:     src.Title,
		NetWorth:  formatMoney(src.NetWorth),
		Money:     formatMoney(src.Money),
		Salary:    formatMoney(src.Salary),
		Inventory: inventory,
	}
}

type Resource struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

func NewResource(src *game.Resource) Resource {
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

	BuyPrice  string `json:"buy_price"`
	SellPrice string `json:"sell_price"`
}

func NewListing(src *game.Listing) Listing {
	return Listing{
		Resource:  NewResource(src.Resource),
		BuyPrice:  formatMoney(src.BuyPrice),
		SellPrice: formatMoney(src.SellPrice),
	}
}

func formatMoney(src float64) string {
	return fmt.Sprintf("$%.2f", src)
}
