package game

// Master list of all resources used in the game. Modifying this map or an individual listing will affect all players.
var stockMasterList = MapMerge(createStockMap(commodityList), createStockMap(equipmentList), createStockMap(landList), createStockMap(employeeList))

// Takes a slice of listings and converts into a map for easy lookup by name.
func createStockMap(src []*Resource) map[string]*Listing {
	listings := map[string]*Listing{}
	for _, v := range src {
		listings[v.Name] = &Listing{
			Resource: v,
		}
	}
	return listings
}
