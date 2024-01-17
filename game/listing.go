package game

const (
	LISTING_FILTER_NAME = "Name"
	LISTING_FILTER_TYPE = "Type"
)

// Filter to be used on market listings.
type ListingFilter struct {
	Property string
	Value    string
}

// Represents a listing on the market for a resource.
type Listing struct {
	Resource *Resource

	BuyPrice  float64
	SellPrice float64
}

// Adjust the market price for this listing.
func (l *Listing) adjustMarketPrice() {
	l.BuyPrice = randFloat64(l.Resource.buyRangeLow, l.Resource.buyRangeHigh)
	l.SellPrice = l.BuyPrice - randFloat64(0.01, l.Resource.sellDelta)
}
