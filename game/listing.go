package game

const (
	LISTING_FILTER_NAME = "Name"
	LISTING_FILTER_TYPE = "Type"
)

// TODO: where can I put these lists so they aren't global and cumbersome?
// Master list of all resources used in the game. Modifying this map or an individual listing will affect all players.
var resourceMap = MapMerge(createListings(commodityList), createListings(equipmentList), createListings(landList), createListings(employeeList))

// Master list of just the commodities.
var commodityMap = createListings(commodityList)

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

// Takes a slice of resources and converts into a map for easy lookup by name.
func createListings(src []*Resource) map[string]*Listing {
	listings := map[string]*Listing{}
	for _, v := range src {
		listings[v.Name] = &Listing{
			Resource: v,
		}
	}
	return listings
}

// Find and retrieve a resource from the master list.
func findResource(name string) *Resource {
	listing, ok := commodityMap[name]
	if !ok {
		return nil
	}
	return listing.Resource
}
