package game

type ResourceType string

const (
	RESOURCE_TYPE_COMMODITY ResourceType = "Commodity"
	RESOURCE_TYPE_EQUIPMENT ResourceType = "Equipment"
	RESOURCE_TYPE_LAND      ResourceType = "Land"
	RESOURCE_TYPE_EMPLOYEE  ResourceType = "Employee"
)

// Represents a singular resource that can be obtained.
type Resource struct {
	Name string
	Type ResourceType

	// If set, this function will be called every world Update. Useful for
	// generating materials or performing some action.
	Update func(player *Player, item *Item)

	BuyRangeLow  float64 // lowest possible buy price for resource
	BuyRangeHigh float64 // highest possible buy price for resource
	SellDelta    float64 // highest potential difference of selling price. Selling price is always lower than buying price

	// If set, this function will be called before this item can be bought. Can
	// be used to check for prereqs to purchasing equipment. If the prereqs are
	// met, can be used to consume items as part of the purchase process.
	prebuy func(player *Player) bool
}

// List of all commodities in the game.
var commodityList = []*Resource{
	{
		Name:         "Limestone",
		Type:         RESOURCE_TYPE_COMMODITY,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Sandstone",
		Type:         RESOURCE_TYPE_COMMODITY,
		BuyRangeLow:  8,
		BuyRangeHigh: 30,
		SellDelta:    7,
	},
	{
		Name:         "Granite",
		Type:         RESOURCE_TYPE_COMMODITY,
		BuyRangeLow:  5,
		BuyRangeHigh: 50,
		SellDelta:    4,
	},
	{
		Name:         "Marble",
		Type:         RESOURCE_TYPE_COMMODITY,
		BuyRangeLow:  10,
		BuyRangeHigh: 100,
		SellDelta:    9,
	},
	{
		Name:         "Quartz",
		Type:         RESOURCE_TYPE_COMMODITY,
		BuyRangeLow:  25,
		BuyRangeHigh: 90,
		SellDelta:    20,
	},
	{
		Name:         "Onyx",
		Type:         RESOURCE_TYPE_COMMODITY,
		BuyRangeLow:  30,
		BuyRangeHigh: 100,
		SellDelta:    20,
	},
	{
		Name:         "Jade",
		Type:         RESOURCE_TYPE_COMMODITY,
		BuyRangeLow:  50,
		BuyRangeHigh: 150,
		SellDelta:    40,
	},
	{
		Name:         "Garnet",
		Type:         RESOURCE_TYPE_COMMODITY,
		BuyRangeLow:  300,
		BuyRangeHigh: 600,
		SellDelta:    200,
	},
	{
		Name:         "Sapphire",
		Type:         RESOURCE_TYPE_COMMODITY,
		BuyRangeLow:  1000,
		BuyRangeHigh: 3000,
		SellDelta:    500,
	},
	{
		Name:         "Ruby",
		Type:         RESOURCE_TYPE_COMMODITY,
		BuyRangeLow:  1000,
		BuyRangeHigh: 3000,
		SellDelta:    500,
	},
	{
		Name:         "Emerald",
		Type:         RESOURCE_TYPE_COMMODITY,
		BuyRangeLow:  1000,
		BuyRangeHigh: 3000,
		SellDelta:    500,
	},
	{
		Name:         "Diamond",
		Type:         RESOURCE_TYPE_COMMODITY,
		BuyRangeLow:  6000,
		BuyRangeHigh: 12000,
		SellDelta:    1000,
	},
	{
		Name:         "Copper",
		Type:         RESOURCE_TYPE_COMMODITY,
		BuyRangeLow:  20,
		BuyRangeHigh: 50,
		SellDelta:    10,
	},
	{
		Name:         "Silver",
		Type:         RESOURCE_TYPE_COMMODITY,
		BuyRangeLow:  250,
		BuyRangeHigh: 750,
		SellDelta:    100,
	},
	{
		Name:         "Gold",
		Type:         RESOURCE_TYPE_COMMODITY,
		BuyRangeLow:  5000,
		BuyRangeHigh: 10000,
		SellDelta:    2000,
	},
	{
		Name:         "Gold Flakes",
		Type:         RESOURCE_TYPE_COMMODITY,
		BuyRangeLow:  13,
		BuyRangeHigh: 17,
		SellDelta:    2,
	},
	{
		Name:         "Platinum",
		Type:         RESOURCE_TYPE_COMMODITY,
		BuyRangeLow:  7500,
		BuyRangeHigh: 15000,
		SellDelta:    2000,
	},
}

// List of all equipment in the game.
var equipmentList = []*Resource{
	// TODO: add equipment
	{
		Name:         "Pickaxe",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
		// TODO: all properties should be merged together into resource, except the buy and sell price. Resources should be defined in the resource file.
	},
	{
		Name:         "Gold Pan",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Sluice Box",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Water Pump",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Map",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Binoculars",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "GPS Unit",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Mine Cart",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Dynamite",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Small Dump Truck",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Medium Dump Truck",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Large Dump Truck",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Ultra-Heavy Dump Truck",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Small Crane",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Medium Crane",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Large Crane",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Ultra-Heavy Crane",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
}

// List of all land in the game.
var landList = []*Resource{
	// TODO: river claims are for gold flakes only, desert claim for ore (copper through platinum), mountain claim for everything else, including diamonds
	{
		Name:         "River Claim - Low Grade",
		Type:         RESOURCE_TYPE_LAND,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "River Claim - High Grade",
		Type:         RESOURCE_TYPE_LAND,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "River Claim - Superior Grade",
		Type:         RESOURCE_TYPE_LAND,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Mountain Claim - Low Grade",
		Type:         RESOURCE_TYPE_LAND,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Mountain Claim - High Grade",
		Type:         RESOURCE_TYPE_LAND,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Mountain Claim - Superior Grade",
		Type:         RESOURCE_TYPE_LAND,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Desert Claim - Low Grade",
		Type:         RESOURCE_TYPE_LAND,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Desert Claim - High Grade",
		Type:         RESOURCE_TYPE_LAND,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Desert Claim - Superior Grade",
		Type:         RESOURCE_TYPE_LAND,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name: "Hydraulic Mine - Low Yield",
		Type: RESOURCE_TYPE_LAND,
		// TODO: update func for hydraulic mine
		// update: func(player *Player, item *Item) {
		// 	player.AddResource(commodityList["Gold Flakes"].Resource, randInt64(1*item.Quantity, 5*item.Quantity))
		// },
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
		prebuy: func(player *Player) bool {
			// TODO: subtract items from player inventory (sluice box, water pump, employees)
			return true
		},
	},
	{
		Name:         "Hydraulic Mine - High Yield",
		Type:         RESOURCE_TYPE_LAND,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Hydraulic Mine - Superior Yield",
		Type:         RESOURCE_TYPE_LAND,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
}

// List of all employees in the game.
var employeeList = []*Resource{
	// TODO: finish filling out employees
	{
		Name:         "Worker",
		Type:         RESOURCE_TYPE_EMPLOYEE,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Surveyor",
		Type:         RESOURCE_TYPE_EMPLOYEE,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Specialist",
		Type:         RESOURCE_TYPE_EMPLOYEE,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
	{
		Name:         "Mining Engineer",
		Type:         RESOURCE_TYPE_EMPLOYEE,
		BuyRangeLow:  5,
		BuyRangeHigh: 20,
		SellDelta:    4,
	},
}
