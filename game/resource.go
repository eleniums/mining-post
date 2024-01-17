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
	Name        string
	Description string
	Type        ResourceType

	// If set, this function will be called every world update. Useful for
	// generating materials or performing some action.
	update func(player *Player, item *Item)

	buyRangeLow  float64 // lowest possible buy price for resource
	buyRangeHigh float64 // highest possible buy price for resource
	sellDelta    float64 // highest potential difference of selling price. Selling price is always lower than buying price

	// If set, this function will be called before this item can be bought. Can
	// be used to check for prereqs to purchasing equipment. If the prereqs are
	// met, can be used to consume items as part of the purchase process.
	prebuy func(player *Player) bool
}

// List of all commodities in the game.
var commodityList = []*Resource{
	{
		Name:         "Limestone",
		Description:  "Limestone is a sedimentary rock composed mainly of calcium carbonate, often used in construction for its durability.",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Sandstone",
		Description:  "Sandstone is a sedimentary rock with a granular texture, commonly employed in construction and known for its porous qualities.",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  8,
		buyRangeHigh: 30,
		sellDelta:    7,
	},
	{
		Name:         "Granite",
		Description:  "Granite is a coarse-grained, igneous rock with a speckled appearance, widely used in construction for its hardness and durability.",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  5,
		buyRangeHigh: 50,
		sellDelta:    4,
	},
	{
		Name:         "Marble",
		Description:  "Marble is a metamorphic rock prized for its smooth texture and distinctive veining, commonly used in sculpture and building materials.",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  10,
		buyRangeHigh: 100,
		sellDelta:    9,
	},
	{
		Name:         "Quartz",
		Description:  "Quartz is a hard, crystalline mineral composed of silicon and oxygen atoms, often used in the production of jewelry and electronic devices for its durability and piezoelectric properties.",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  25,
		buyRangeHigh: 90,
		sellDelta:    20,
	},
	{
		Name:         "Onyx",
		Description:  "Onyx is a type of cryptocrystalline quartz known for its smooth, banded layers of alternating colors, often used in jewelry and decorative items.",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  30,
		buyRangeHigh: 100,
		sellDelta:    20,
	},
	{
		Name:         "Jade",
		Description:  "Jade is a dense, ornamental mineral prized for its rich green color, often used in traditional art and jewelry.",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  50,
		buyRangeHigh: 150,
		sellDelta:    40,
	},
	{
		Name:         "Garnet",
		Description:  "Garnet is a group of minerals with deep red to vibrant green hues, commonly used as gemstones and abrasives.",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  300,
		buyRangeHigh: 600,
		sellDelta:    200,
	},
	{
		Name:         "Sapphire",
		Description:  "Sapphire is a precious gemstone, typically blue in color, composed of corundum and valued for its hardness and luster.",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  1000,
		buyRangeHigh: 3000,
		sellDelta:    500,
	},
	{
		Name:         "Ruby",
		Description:  "Ruby is a red precious gemstone, a variety of corundum, prized for its vibrant color and considered one of the most precious gems.",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  1000,
		buyRangeHigh: 3000,
		sellDelta:    500,
	},
	{
		Name:         "Emerald",
		Description:  "Emerald is a green precious gemstone, a variety of the mineral beryl, celebrated for its vivid color and historical significance.",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  1000,
		buyRangeHigh: 3000,
		sellDelta:    500,
	},
	{
		Name:         "Diamond",
		Description:  "Diamond is a brilliant and exceptionally hard precious gemstone, composed of carbon atoms arranged in a crystal lattice, prized for its rarity and used in various luxury applications.",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  6000,
		buyRangeHigh: 12000,
		sellDelta:    1000,
	},
	{
		Name:         "Copper",
		Description:  "Copper is a reddish-brown metal known for its conductivity and malleability.",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  20,
		buyRangeHigh: 50,
		sellDelta:    10,
	},
	{
		Name:         "Silver",
		Description:  "Silver is a lustrous, precious metal admired for its conductivity and versatile use in jewelry and industry.",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  250,
		buyRangeHigh: 750,
		sellDelta:    100,
	},
	{
		Name:         "Gold",
		Description:  "Gold is a prized, yellow metal celebrated for its rarity, beauty, and historical significance in various cultural and economic contexts.",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  5000,
		buyRangeHigh: 10000,
		sellDelta:    2000,
	},
	{
		Name:         "Gold Flakes",
		Description:  "Gold flakes are tiny, flat, and thin pieces of gold that are small in size and often resemble flakes of paint or glitter.",
		Type:         RESOURCE_TYPE_COMMODITY,
		buyRangeLow:  13,
		buyRangeHigh: 17,
		sellDelta:    2,
	},
	{
		Name:         "Platinum",
		Description:  "Platinum is a dense, silvery-white metal valued for its durability and use in jewelry and catalytic converters.",
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
		Description:  "",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
		// TODO: all properties should be merged together into resource, except the buy and sell price. Resources should be defined in the resource file.
	},
	{
		Name:         "Gold Pan",
		Description:  "",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Sluice Box",
		Description:  "",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Water Pump",
		Description:  "",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Map",
		Description:  "",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Binoculars",
		Description:  "",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "GPS Unit",
		Description:  "",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Mine Cart",
		Description:  "",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Dynamite",
		Description:  "",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Small Dump Truck",
		Description:  "",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Medium Dump Truck",
		Description:  "",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Large Dump Truck",
		Description:  "",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Ultra-Heavy Dump Truck",
		Description:  "",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Small Crane",
		Description:  "",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Medium Crane",
		Description:  "",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Large Crane",
		Description:  "",
		Type:         RESOURCE_TYPE_EQUIPMENT,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Ultra-Heavy Crane",
		Description:  "",
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
		Description:  "",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "River Claim - High Grade",
		Description:  "",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "River Claim - Superior Grade",
		Description:  "",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Mountain Claim - Low Grade",
		Description:  "",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Mountain Claim - High Grade",
		Description:  "",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Mountain Claim - Superior Grade",
		Description:  "",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Desert Claim - Low Grade",
		Description:  "",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Desert Claim - High Grade",
		Description:  "",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Desert Claim - Superior Grade",
		Description:  "",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:        "Hydraulic Mine - Low Yield",
		Description: "A hydraulic mine is a mining method that uses high-pressure water jets to dislodge soil and gravel for the extraction of gold.",
		Type:        RESOURCE_TYPE_LAND,
		// TODO: update func for hydraulic mine
		// update: func(player *Player, item *Item) {
		// 	player.AddResource(commodityList["Gold Flakes"].Resource, randInt64(1*item.Quantity, 5*item.Quantity))
		// },
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
		prebuy: func(player *Player) bool {
			// TODO: subtract items from player inventory (sluice box, water pump, employees)
			return true
		},
	},
	{
		Name:         "Hydraulic Mine - High Yield",
		Description:  "A hydraulic mine is a mining method that uses high-pressure water jets to dislodge soil and gravel for the extraction of gold.",
		Type:         RESOURCE_TYPE_LAND,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Hydraulic Mine - Superior Yield",
		Description:  "A hydraulic mine is a mining method that uses high-pressure water jets to dislodge soil and gravel for the extraction of gold.",
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
		Description:  "",
		Type:         RESOURCE_TYPE_EMPLOYEE,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Surveyor",
		Description:  "",
		Type:         RESOURCE_TYPE_EMPLOYEE,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Specialist",
		Description:  "",
		Type:         RESOURCE_TYPE_EMPLOYEE,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
	{
		Name:         "Mining Engineer",
		Description:  "",
		Type:         RESOURCE_TYPE_EMPLOYEE,
		buyRangeLow:  5,
		buyRangeHigh: 20,
		sellDelta:    4,
	},
}
