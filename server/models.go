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

type Item struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Quantity int64  `json:"quantity"`
}

func NewItem(src *game.Item) Item {
	return Item{
		Name:     src.Resource.Name,
		Type:     string(src.Resource.Type),
		Quantity: src.Quantity,
	}
}

type Listing struct {
	Name          string         `json:"name"`
	Type          string         `json:"type"`
	BuyPrice      string         `json:"buy_price"`
	SellPrice     string         `json:"sell_price"`
	Prerequisites []Prerequisite `json:"prerequisites,omitempty"`
}

func NewListing(src *game.Listing) Listing {
	prerequisites := make([]Prerequisite, len(src.Resource.Prerequisites))
	for i, v := range src.Resource.Prerequisites {
		prerequisites[i] = Prerequisite{
			Name:     v.Name,
			Quantity: v.Quantity,
		}
	}

	return Listing{
		Name:          src.Resource.Name,
		Type:          string(src.Resource.Type),
		BuyPrice:      formatMoney(src.BuyPrice),
		SellPrice:     formatMoney(src.SellPrice),
		Prerequisites: prerequisites,
	}
}

type Prerequisite struct {
	Name     string `json:"name"`
	Quantity int64  `json:"quantity"`
}

func formatMoney(src float64) string {
	// TODO: add commas to money
	return fmt.Sprintf("$%.2f", src)
}
