package game

const (
	LISTING_FILTER_NAME = "Name"
	LISTING_FILTER_TYPE = "Type"
)

type ListingFilter struct {
	Property string
	Value    string
}

// Represents a listing on the market.
type Listing struct {
	Resource *Resource

	BuyPrice  float64
	SellPrice float64

	buyRangeLow  float64 // lowest possible buy price for resource
	buyRangeHigh float64 // highest possible buy price for resource
	sellDelta    float64 // highest potential difference of selling price. Selling price is always lower than buying price

	// If set, this function will be called before this item can be bought. Can
	// be used to check for prereqs to purchasing equipment. If the prereqs are
	// met, can be used to consume items as part of the purchase process.
	prebuy func(player *Player) bool

	// If set, this function will be called before this item can be sold. Can be
	// used to check for prereqs to selling equipment. If the prereqs are met,
	// can be used to add inventory back as part of the selling process.
	presell func(player *Player) bool
}

// Adjust the market price for this listing.
func (l *Listing) adjustMarketPrice() {
	l.BuyPrice = randFloat64(l.buyRangeLow, l.buyRangeHigh)
	l.SellPrice = l.BuyPrice - randFloat64(0.01, l.sellDelta)
}
