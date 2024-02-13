package game

// Master list of all resources used in the game. Modifying this map or an individual listing will affect all players.
var resourceMap = MapMerge(createListings(commodityList), createListings(equipmentList), createListings(landList), createListings(employeeList))

// List of all commodities in the game.
var commodityList = []*Resource{
	{
		Name:      "Limestone",
		Type:      RESOURCE_TYPE_COMMODITY,
		buyLow:    5,
		buyHigh:   20,
		sellDelta: 4,
	},
	{
		Name:      "Sandstone",
		Type:      RESOURCE_TYPE_COMMODITY,
		buyLow:    8,
		buyHigh:   30,
		sellDelta: 7,
	},
	{
		Name:      "Granite",
		Type:      RESOURCE_TYPE_COMMODITY,
		buyLow:    5,
		buyHigh:   50,
		sellDelta: 4,
	},
	{
		Name:      "Quartz",
		Type:      RESOURCE_TYPE_COMMODITY,
		buyLow:    25,
		buyHigh:   90,
		sellDelta: 20,
	},
	{
		Name:      "Marble",
		Type:      RESOURCE_TYPE_COMMODITY,
		buyLow:    10,
		buyHigh:   100,
		sellDelta: 9,
	},
	{
		Name:      "Onyx",
		Type:      RESOURCE_TYPE_COMMODITY,
		buyLow:    30,
		buyHigh:   100,
		sellDelta: 20,
	},
	{
		Name:      "Jade",
		Type:      RESOURCE_TYPE_COMMODITY,
		buyLow:    50,
		buyHigh:   150,
		sellDelta: 40,
	},
	{
		Name:      "Garnet",
		Type:      RESOURCE_TYPE_COMMODITY,
		buyLow:    300,
		buyHigh:   600,
		sellDelta: 200,
	},
	{
		Name:      "Sapphire",
		Type:      RESOURCE_TYPE_COMMODITY,
		buyLow:    1000,
		buyHigh:   3000,
		sellDelta: 500,
	},
	{
		Name:      "Ruby",
		Type:      RESOURCE_TYPE_COMMODITY,
		buyLow:    1000,
		buyHigh:   3000,
		sellDelta: 500,
	},
	{
		Name:      "Emerald",
		Type:      RESOURCE_TYPE_COMMODITY,
		buyLow:    1000,
		buyHigh:   3000,
		sellDelta: 500,
	},
	{
		Name:      "Diamond",
		Type:      RESOURCE_TYPE_COMMODITY,
		buyLow:    6000,
		buyHigh:   12000,
		sellDelta: 1000,
	},
	{
		Name:      "Copper",
		Type:      RESOURCE_TYPE_COMMODITY,
		buyLow:    20,
		buyHigh:   50,
		sellDelta: 10,
	},
	{
		Name:      "Silver",
		Type:      RESOURCE_TYPE_COMMODITY,
		buyLow:    250,
		buyHigh:   750,
		sellDelta: 100,
	},
	{
		Name:      "Gold",
		Type:      RESOURCE_TYPE_COMMODITY,
		buyLow:    5000,
		buyHigh:   10000,
		sellDelta: 2000,
	},
	{
		Name:      "Gold Flakes",
		Type:      RESOURCE_TYPE_COMMODITY,
		buyLow:    13,
		buyHigh:   17,
		sellDelta: 2,
	},
	{
		Name:      "Platinum",
		Type:      RESOURCE_TYPE_COMMODITY,
		buyLow:    7500,
		buyHigh:   15000,
		sellDelta: 2000,
	},
}

// List of all equipment in the game.
var equipmentList = []*Resource{
	{
		Name:      "Pickaxe",
		Type:      RESOURCE_TYPE_EQUIPMENT,
		buyLow:    10,
		buyHigh:   10,
		sellDelta: 0,
	},
	{
		Name:      "Gold Pan",
		Type:      RESOURCE_TYPE_EQUIPMENT,
		buyLow:    15,
		buyHigh:   15,
		sellDelta: 0,
	},
	{
		Name:      "Sluice Box",
		Type:      RESOURCE_TYPE_EQUIPMENT,
		buyLow:    200,
		buyHigh:   200,
		sellDelta: 0,
	},
	{
		Name:      "Water Pump",
		Type:      RESOURCE_TYPE_EQUIPMENT,
		buyLow:    1000,
		buyHigh:   1000,
		sellDelta: 0,
	},
	{
		Name:      "Map",
		Type:      RESOURCE_TYPE_EQUIPMENT,
		buyLow:    5,
		buyHigh:   5,
		sellDelta: 0,
	},
	{
		Name:      "Binoculars",
		Type:      RESOURCE_TYPE_EQUIPMENT,
		buyLow:    20,
		buyHigh:   20,
		sellDelta: 0,
	},
	{
		Name:      "GPS Unit",
		Type:      RESOURCE_TYPE_EQUIPMENT,
		buyLow:    100,
		buyHigh:   100,
		sellDelta: 0,
	},
	{
		Name:      "Mine Cart",
		Type:      RESOURCE_TYPE_EQUIPMENT,
		buyLow:    500,
		buyHigh:   500,
		sellDelta: 0,
	},
	{
		Name:      "Dynamite",
		Type:      RESOURCE_TYPE_EQUIPMENT,
		buyLow:    20,
		buyHigh:   20,
		sellDelta: 0,
	},
	{
		Name:      "Small Dump Truck",
		Type:      RESOURCE_TYPE_EQUIPMENT,
		buyLow:    60_000,
		buyHigh:   60_000,
		sellDelta: 0,
	},
	{
		Name:      "Large Dump Truck",
		Type:      RESOURCE_TYPE_EQUIPMENT,
		buyLow:    500_000,
		buyHigh:   500_000,
		sellDelta: 0,
	},
	{
		Name:      "Ultra-Heavy Dump Truck",
		Type:      RESOURCE_TYPE_EQUIPMENT,
		buyLow:    1_500_000,
		buyHigh:   1_500_000,
		sellDelta: 0,
	},
	{
		Name:      "Small Crane",
		Type:      RESOURCE_TYPE_EQUIPMENT,
		buyLow:    80_000,
		buyHigh:   80_000,
		sellDelta: 0,
	},
	{
		Name:      "Large Crane",
		Type:      RESOURCE_TYPE_EQUIPMENT,
		buyLow:    650_000,
		buyHigh:   650_000,
		sellDelta: 0,
	},
	{
		Name:      "Ultra-Heavy Crane",
		Type:      RESOURCE_TYPE_EQUIPMENT,
		buyLow:    2_000_000,
		buyHigh:   2_000_000,
		sellDelta: 0,
	},
}

// List of all land in the game.
var landList = []*Resource{
	{
		Name:      "Mountain Claim - Low Grade",
		Type:      RESOURCE_TYPE_LAND,
		buyLow:    5,
		buyHigh:   20,
		sellDelta: 0,
	},
	{
		Name:      "Mountain Claim - High Grade",
		Type:      RESOURCE_TYPE_LAND,
		buyLow:    5,
		buyHigh:   20,
		sellDelta: 0,
	},
	{
		Name:      "Mountain Claim - Superior Grade",
		Type:      RESOURCE_TYPE_LAND,
		buyLow:    5,
		buyHigh:   20,
		sellDelta: 0,
	},
	{
		Name:      "Mountain Mine - Low Yield",
		Type:      RESOURCE_TYPE_LAND,
		buyLow:    300_000,
		buyHigh:   300_000,
		sellDelta: 0,
		Prerequisites: []Prerequisite{
			{Name: "Mountain Claim - Low Grade", Quantity: 1},
			{Name: "Small Dump Truck", Quantity: 1},
			{Name: "Small Crane", Quantity: 1},
			{Name: "Worker", Quantity: 3},
		},
		Loot: LootTable{
			{Name: "Limestone", Weight: 11, CountLow: 5, CountHigh: 22}, // max value: 22 * 20 = 440
			{Name: "Sandstone", Weight: 11, CountLow: 5, CountHigh: 15}, // max value: 15 * 30 = 450
			{Name: "Granite", Weight: 11, CountLow: 4, CountHigh: 9},    // max value: 9 * 50 = 450
			{Name: "Quartz", Weight: 11, CountLow: 4, CountHigh: 6},     // max value: 6 * 90 = 540
			{Name: "Marble", Weight: 10, CountLow: 2, CountHigh: 8},     // max value: 8 * 100 = 800
			{Name: "Onyx", Weight: 10, CountLow: 2, CountHigh: 8},       // max value: 8 * 100 = 800
			{Name: "Jade", Weight: 10, CountLow: 2, CountHigh: 8},       // max value: 8 * 150 = 1200
			{Name: "Garnet", Weight: 10, CountLow: 2, CountHigh: 4},     // max value: 4 * 600 = 2400
			{Name: "Sapphire", Weight: 5, CountLow: 1, CountHigh: 2},    // max value: 2 * 3000 = 6000
			{Name: "Ruby", Weight: 5, CountLow: 1, CountHigh: 2},        // max value: 2 * 3000 = 6000
			{Name: "Emerald", Weight: 5, CountLow: 1, CountHigh: 2},     // max value: 2 * 3000 = 6000
			{Name: "Diamond", Weight: 1, CountLow: 1, CountHigh: 1},     // max value: 1 * 12000 = 12000
		},
	},
	{
		Name:      "Mountain Mine - High Yield",
		Type:      RESOURCE_TYPE_LAND,
		buyLow:    1_500_000,
		buyHigh:   1_500_000,
		sellDelta: 0,
		Prerequisites: []Prerequisite{
			{Name: "Mountain Claim - High Grade", Quantity: 1},
			{Name: "Large Dump Truck", Quantity: 2},
			{Name: "Large Crane", Quantity: 1},
			{Name: "Worker", Quantity: 4},
			{Name: "Specialist", Quantity: 1},
		},
		Loot: LootTable{
			{Name: "Limestone", Weight: 3, CountLow: 20, CountHigh: 50}, // max value: 50 * 20 = 1000
			{Name: "Sandstone", Weight: 3, CountLow: 15, CountHigh: 35}, // max value: 35 * 30 = 1050
			{Name: "Granite", Weight: 3, CountLow: 10, CountHigh: 22},   // max value: 22 * 50 = 1100
			{Name: "Quartz", Weight: 3, CountLow: 8, CountHigh: 15},     // max value: 15 * 90 = 1350
			{Name: "Marble", Weight: 11, CountLow: 10, CountHigh: 20},   // max value: 20 * 100 = 2000
			{Name: "Onyx", Weight: 11, CountLow: 10, CountHigh: 20},     // max value: 20 * 100 = 2000
			{Name: "Jade", Weight: 11, CountLow: 10, CountHigh: 20},     // max value: 20 * 150 = 3000
			{Name: "Garnet", Weight: 11, CountLow: 5, CountHigh: 10},    // max value: 10 * 600 = 6000
			{Name: "Sapphire", Weight: 13, CountLow: 3, CountHigh: 5},   // max value: 5 * 3000 = 15000
			{Name: "Ruby", Weight: 13, CountLow: 3, CountHigh: 5},       // max value: 5 * 3000 = 15000
			{Name: "Emerald", Weight: 13, CountLow: 3, CountHigh: 5},    // max value: 5 * 3000 = 15000
			{Name: "Diamond", Weight: 5, CountLow: 1, CountHigh: 2},     // max value: 2 * 12000 = 24000
		},
	},
	// TODO: update mountain mine values
	{
		Name:      "Mountain Mine - Superior Yield",
		Type:      RESOURCE_TYPE_LAND,
		buyLow:    6_000_000,
		buyHigh:   6_000_000,
		sellDelta: 0,
		Prerequisites: []Prerequisite{
			{Name: "Mountain Claim - Superior Grade", Quantity: 1},
			{Name: "Ultra-Heavy Dump Truck", Quantity: 2},
			{Name: "Ultra-Heavy Crane", Quantity: 1},
			{Name: "Worker", Quantity: 6},
			{Name: "Specialist", Quantity: 3},
			{Name: "Mining Engineer", Quantity: 1},
		},
		Loot: LootTable{
			{Name: "Limestone", Weight: 50, CountLow: 20, CountHigh: 50}, // max value: 0 * 20 = 0
			{Name: "Sandstone", Weight: 50, CountLow: 20, CountHigh: 50}, // max value: 0 * 30 = 0
			{Name: "Granite", Weight: 50, CountLow: 20, CountHigh: 50},   // max value: 0 * 50 = 0
			{Name: "Quartz", Weight: 50, CountLow: 20, CountHigh: 50},    // max value: 0 * 90 = 0
			{Name: "Marble", Weight: 50, CountLow: 20, CountHigh: 50},    // max value: 0 * 100 = 0
			{Name: "Onyx", Weight: 50, CountLow: 20, CountHigh: 50},      // max value: 0 * 100 = 0
			{Name: "Jade", Weight: 50, CountLow: 20, CountHigh: 50},      // max value: 0 * 150 = 0
			{Name: "Garnet", Weight: 50, CountLow: 20, CountHigh: 50},    // max value: 0 * 600 = 0
			{Name: "Sapphire", Weight: 50, CountLow: 20, CountHigh: 50},  // max value: 0 * 3000 = 0
			{Name: "Ruby", Weight: 50, CountLow: 20, CountHigh: 50},      // max value: 0 * 3000 = 0
			{Name: "Emerald", Weight: 50, CountLow: 20, CountHigh: 50},   // max value: 0 * 3000 = 0
			{Name: "Diamond", Weight: 50, CountLow: 20, CountHigh: 50},   // max value: 0 * 12000 = 0
		},
	},
	{
		Name:      "Desert Claim - Low Grade",
		Type:      RESOURCE_TYPE_LAND,
		buyLow:    10_000,
		buyHigh:   10_000,
		sellDelta: 0,
	},
	{
		Name:      "Desert Claim - High Grade",
		Type:      RESOURCE_TYPE_LAND,
		buyLow:    100_000,
		buyHigh:   100_000,
		sellDelta: 0,
	},
	{
		Name:      "Desert Claim - Superior Grade",
		Type:      RESOURCE_TYPE_LAND,
		buyLow:    1_000_000,
		buyHigh:   1_000_000,
		sellDelta: 0,
	},
	{
		Name:      "Desert Mine - Low Yield",
		Type:      RESOURCE_TYPE_LAND,
		buyLow:    600_000,
		buyHigh:   600_000,
		sellDelta: 0,
		Prerequisites: []Prerequisite{
			{Name: "Desert Claim - Low Grade", Quantity: 1},
			{Name: "Small Dump Truck", Quantity: 1},
			{Name: "Small Crane", Quantity: 1},
			{Name: "Worker", Quantity: 3},
		},
		Loot: LootTable{
			{Name: "Copper", Weight: 50, CountLow: 20, CountHigh: 50}, // max value: 50 * 50 = 2500
			{Name: "Silver", Weight: 35, CountLow: 4, CountHigh: 10},  // max value: 750 * 8 = 7500
			{Name: "Gold", Weight: 14, CountLow: 1, CountHigh: 3},     // max value: 10000 * 3 = 30000
			{Name: "Platinum", Weight: 1, CountLow: 1, CountHigh: 3},  // max value: 15000 * 3 = 45000
		},
	},
	{
		Name:      "Desert Mine - High Yield",
		Type:      RESOURCE_TYPE_LAND,
		buyLow:    2_500_000,
		buyHigh:   2_500_000,
		sellDelta: 0,
		Prerequisites: []Prerequisite{
			{Name: "Desert Claim - High Grade", Quantity: 1},
			{Name: "Large Dump Truck", Quantity: 2},
			{Name: "Large Crane", Quantity: 1},
			{Name: "Worker", Quantity: 4},
			{Name: "Specialist", Quantity: 1},
		},
		Loot: LootTable{
			{Name: "Copper", Weight: 20, CountLow: 50, CountHigh: 100}, // max value: 50 * 100 = 5000
			{Name: "Silver", Weight: 50, CountLow: 15, CountHigh: 30},  // max value: 750 * 30 = 22500
			{Name: "Gold", Weight: 20, CountLow: 5, CountHigh: 10},     // max value: 10000 * 10 = 100000
			{Name: "Platinum", Weight: 10, CountLow: 5, CountHigh: 10}, // max value: 15000 * 10 = 150000
		},
	},
	{
		Name:      "Desert Mine - Superior Yield",
		Type:      RESOURCE_TYPE_LAND,
		buyLow:    10_000_000,
		buyHigh:   10_000_000,
		sellDelta: 0,
		Prerequisites: []Prerequisite{
			{Name: "Desert Claim - Superior Grade", Quantity: 1},
			{Name: "Ultra-Heavy Dump Truck", Quantity: 2},
			{Name: "Ultra-Heavy Crane", Quantity: 1},
			{Name: "Worker", Quantity: 6},
			{Name: "Specialist", Quantity: 3},
			{Name: "Mining Engineer", Quantity: 1},
		},
		Loot: LootTable{
			{Name: "Copper", Weight: 14, CountLow: 150, CountHigh: 300}, // max value: 50 * 300 = 15000
			{Name: "Silver", Weight: 16, CountLow: 40, CountHigh: 70},   // max value: 750 * 70 = 52500
			{Name: "Gold", Weight: 35, CountLow: 20, CountHigh: 50},     // max value: 10000 * 50 = 500000
			{Name: "Platinum", Weight: 35, CountLow: 20, CountHigh: 50}, // max value: 15000 * 50 = 750000
		},
	},
	{
		Name:      "River Claim - Low Grade",
		Type:      RESOURCE_TYPE_LAND,
		buyLow:    1_000,
		buyHigh:   1_000,
		sellDelta: 0,
	},
	{
		Name:      "River Claim - High Grade",
		Type:      RESOURCE_TYPE_LAND,
		buyLow:    10_000,
		buyHigh:   10_000,
		sellDelta: 0,
	},
	{
		Name:      "River Claim - Superior Grade",
		Type:      RESOURCE_TYPE_LAND,
		buyLow:    100_000,
		buyHigh:   100_000,
		sellDelta: 0,
	},
	{
		Name:      "Hydraulic Mine - Low Yield",
		Type:      RESOURCE_TYPE_LAND,
		buyLow:    50_000,
		buyHigh:   50_000,
		sellDelta: 0,
		Prerequisites: []Prerequisite{
			{Name: "River Claim - Low Grade", Quantity: 1},
			{Name: "Sluice Box", Quantity: 1},
			{Name: "Water Pump", Quantity: 1},
			{Name: "Worker", Quantity: 1},
		},
		Loot: LootTable{
			{Name: "Gold Flakes", Weight: 100, CountLow: 1, CountHigh: 5}, // max value: 17 * 5 = 85
		},
	},
	{
		Name:      "Hydraulic Mine - High Yield",
		Type:      RESOURCE_TYPE_LAND,
		buyLow:    450_000,
		buyHigh:   450_000,
		sellDelta: 0,
		Prerequisites: []Prerequisite{
			{Name: "River Claim - High Grade", Quantity: 1},
			{Name: "Sluice Box", Quantity: 3},
			{Name: "Water Pump", Quantity: 3},
			{Name: "Worker", Quantity: 2},
		},
		Loot: LootTable{
			{Name: "Gold Flakes", Weight: 60, CountLow: 10, CountHigh: 20}, // max value: 17 * 20 = 340
			{Name: "Gold Flakes", Weight: 29, CountLow: 30, CountHigh: 50}, // max value: 17 * 50 = 850
			{Name: "Gold", Weight: 1, CountLow: 1, CountHigh: 1},           // max value: 10000 * 1 = 10000
		},
	},
	{
		Name:      "Hydraulic Mine - Superior Yield",
		Type:      RESOURCE_TYPE_LAND,
		buyLow:    950_000,
		buyHigh:   950_000,
		sellDelta: 0,
		Prerequisites: []Prerequisite{
			{Name: "River Claim - Superior Grade", Quantity: 1},
			{Name: "Sluice Box", Quantity: 6},
			{Name: "Water Pump", Quantity: 6},
			{Name: "Worker", Quantity: 3},
			{Name: "Specialist", Quantity: 1},
		},
		Loot: LootTable{
			{Name: "Gold Flakes", Weight: 50, CountLow: 30, CountHigh: 50},  // max value: 17 * 50 = 850
			{Name: "Gold Flakes", Weight: 35, CountLow: 60, CountHigh: 100}, // max value: 17 * 100 = 1700
			{Name: "Gold", Weight: 15, CountLow: 1, CountHigh: 3},           // max value: 10000 * 3 = 30000
		},
	},
}

// List of all employees in the game.
var employeeList = []*Resource{
	{
		Name:      "Worker",
		Type:      RESOURCE_TYPE_EMPLOYEE,
		buyLow:    30_000,
		buyHigh:   30_000,
		sellDelta: 0,
	},
	{
		Name:      "Surveyor",
		Type:      RESOURCE_TYPE_EMPLOYEE,
		buyLow:    45_000,
		buyHigh:   45_000,
		sellDelta: 0,
	},
	{
		Name:      "Specialist",
		Type:      RESOURCE_TYPE_EMPLOYEE,
		buyLow:    85_000,
		buyHigh:   85_000,
		sellDelta: 0,
	},
	{
		Name:      "Mining Engineer",
		Type:      RESOURCE_TYPE_EMPLOYEE,
		buyLow:    150_000,
		buyHigh:   150_000,
		sellDelta: 0,
	},
}
