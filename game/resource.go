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

	// Prerequisites to purchasing a resource. If the prerequisites are all met,
	// they will be consumed from the player's inventory as part of the purchase
	// process.
	Prerequisites []Prerequisite

	// If set, this function will be called every world update. Useful for
	// generating materials or performing some action.
	update func(player *Player, item *Item)

	buyRangeLow  float64 // lowest possible buy price for resource
	buyRangeHigh float64 // highest possible buy price for resource
	sellDelta    float64 // highest potential difference of selling price. Selling price is always lower than buying price
}

// Prerequisite for purchasing a resource. Includes name of resource and amount to subtract from player inventory.
type Prerequisite struct {
	Name     string
	Quantity int64
}

// List of all commodities in the game.
var commodityList = []*Resource{
	{
		Name:         "Limestone",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Sandstone",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  8,
		buyRangeHigh: 30,
		sellDelta:    7,
	},
	{
		Name:         "Granite",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  5,
		buyRangeHigh: 50,
		sellDelta:    4,
	},
	{
		Name:         "Marble",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  10,
		buyRangeHigh: 100,
		sellDelta:    9,
	},
	{
		Name:         "Quartz",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  25,
		buyRangeHigh: 90,
		sellDelta:    20,
	},
	{
		Name:         "Onyx",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  30,
		buyRangeHigh: 100,
		sellDelta:    20,
	},
	{
		Name:         "Jade",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  50,
		buyRangeHigh: 150,
		sellDelta:    40,
	},
	{
		Name:         "Garnet",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  300,
		buyRangeHigh: 600,
		sellDelta:    200,
	},
	{
		Name:         "Sapphire",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  1000,
		buyRangeHigh: 3000,
		sellDelta:    500,
	},
	{
		Name:         "Ruby",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  1000,
		buyRangeHigh: 3000,
		sellDelta:    500,
	},
	{
		Name:         "Emerald",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  1000,
		buyRangeHigh: 3000,
		sellDelta:    500,
	},
	{
		Name:         "Diamond",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  6000,
		buyRangeHigh: 12000,
		sellDelta:    1000,
	},
	{
		Name:         "Copper",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  20,
		buyRangeHigh: 50,
		sellDelta:    10,
	},
	{
		Name:         "Silver",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  250,
		buyRangeHigh: 750,
		sellDelta:    100,
	},
	{
		Name:         "Gold",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  5000,
		buyRangeHigh: 10000,
		sellDelta:    2000,
	},
	{
		Name:         "Gold Flakes",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  13,
		buyRangeHigh: 17,
		sellDelta:    2,
	},
	{
		Name:         "Platinum",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  7500,
		buyRangeHigh: 15000,
		sellDelta:    2000,
	},
}

// List of all equipment in the game.
var equipmentList = []*Resource{
	// TODO: add equipment
	{
		Name:         "Pickaxe",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
		// TODO: all properties should be merged together into resource, except the buy and sell price. Resources should be defined in the resource file.
	},
	{
		Name:         "Gold Pan",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Sluice Box",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Water Pump",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Map",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Binoculars",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "GPS Unit",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Mine Cart",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Dynamite",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Small Dump Truck",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Medium Dump Truck",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Large Dump Truck",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Ultra-Heavy Dump Truck",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Small Crane",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Medium Crane",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Large Crane",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Ultra-Heavy Crane",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
}

// List of all land in the game.
var landList = []*Resource{
	// TODO: river claims are for gold flakes only, desert claim for ore (copper through platinum), mountain claim for everything else, including diamonds
	{
		Name:         "River Claim - Low Grade",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "River Claim - High Grade",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "River Claim - Superior Grade",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Mountain Claim - Low Grade",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Mountain Claim - High Grade",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Mountain Claim - Superior Grade",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Desert Claim - Low Grade",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Desert Claim - High Grade",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Desert Claim - Superior Grade",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name: "Hydraulic Mine - Low Yield",
		Type: RESOURCE_TYPE_LAND,
		// TODO: update func for hydraulic mine
		// update: func(player *Player, item *Item) {
		// 	player.AddResource(commodityList["Gold Flakes"].Resource, randInt64(1*item.Quantity, 5*item.Quantity))
		// },
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
		// Prerequisites: func(player *Player) bool {
		// TODO: subtract items from player inventory (sluice box, water pump, employees)
		// 	return true
		// },
	},
	{
		Name:         "Hydraulic Mine - High Yield",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Hydraulic Mine - Superior Yield",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
}

// List of all employees in the game.
var employeeList = []*Resource{
	// TODO: finish filling out employees
	{
		Name:         "Worker",
		Type:         RESOURCE_TYPE_EMPLOYEE,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Surveyor",
		Type:         RESOURCE_TYPE_EMPLOYEE,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Specialist",
		Type:         RESOURCE_TYPE_EMPLOYEE,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Mining Engineer",
		Type:         RESOURCE_TYPE_EMPLOYEE,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
}
