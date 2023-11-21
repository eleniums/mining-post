package game

// Get default list of all stock and prices.
func getInitialStock() []*Listing {
	return []*Listing{
		{
			Resource: Resource{
				Name:        "Limestone",
				Description: "Limestone is a sedimentary rock composed mainly of calcium carbonate, often used in construction for its durability.",
			},
			buyRangeLow:  5,
			buyRangeHigh: 20,
			sellDelta:    4,
		},
		{
			Resource: Resource{
				Name:        "Sandstone",
				Description: "Sandstone is a sedimentary rock with a granular texture, commonly employed in construction and known for its porous qualities.",
			},
			buyRangeLow:  8,
			buyRangeHigh: 30,
			sellDelta:    7,
		},
		{
			Resource: Resource{
				Name:        "Granite",
				Description: "Granite is a coarse-grained, igneous rock with a speckled appearance, widely used in construction for its hardness and durability.",
			},
			buyRangeLow:  5,
			buyRangeHigh: 50,
			sellDelta:    4,
		},
		{
			Resource: Resource{
				Name:        "Marble",
				Description: "Marble is a metamorphic rock prized for its smooth texture and distinctive veining, commonly used in sculpture and building materials.",
			},
			buyRangeLow:  10,
			buyRangeHigh: 100,
			sellDelta:    9,
		},
		{
			Resource: Resource{
				Name:        "Quartz",
				Description: "Quartz is a hard, crystalline mineral composed of silicon and oxygen atoms, often used in the production of jewelry and electronic devices for its durability and piezoelectric properties.",
			},
			buyRangeLow:  25,
			buyRangeHigh: 90,
			sellDelta:    20,
		},
		{
			Resource: Resource{
				Name:        "Onyx",
				Description: "Onyx is a type of cryptocrystalline quartz known for its smooth, banded layers of alternating colors, often used in jewelry and decorative items.",
			},
			buyRangeLow:  30,
			buyRangeHigh: 100,
			sellDelta:    20,
		},
		{
			Resource: Resource{
				Name:        "Jade",
				Description: "Jade is a dense, ornamental mineral prized for its rich green color, often used in traditional art and jewelry.",
			},
			buyRangeLow:  50,
			buyRangeHigh: 150,
			sellDelta:    40,
		},
		{
			Resource: Resource{
				Name:        "Garnet",
				Description: "Garnet is a group of minerals with deep red to vibrant green hues, commonly used as gemstones and abrasives.",
			},
			buyRangeLow:  300,
			buyRangeHigh: 600,
			sellDelta:    200,
		},
		{
			Resource: Resource{
				Name:        "Sapphire",
				Description: "Sapphire is a precious gemstone, typically blue in color, composed of corundum and valued for its hardness and luster.",
			},
			buyRangeLow:  1000,
			buyRangeHigh: 3000,
			sellDelta:    500,
		},
		{
			Resource: Resource{
				Name:        "Ruby",
				Description: "Ruby is a red precious gemstone, a variety of corundum, prized for its vibrant color and considered one of the most precious gems.",
			},
			buyRangeLow:  1000,
			buyRangeHigh: 3000,
			sellDelta:    500,
		},
		{
			Resource: Resource{
				Name:        "Emerald",
				Description: "Emerald is a green precious gemstone, a variety of the mineral beryl, celebrated for its vivid color and historical significance.",
			},
			buyRangeLow:  1000,
			buyRangeHigh: 3000,
			sellDelta:    500,
		},
		{
			Resource: Resource{
				Name:        "Diamond",
				Description: "Diamond is a brilliant and exceptionally hard precious gemstone, composed of carbon atoms arranged in a crystal lattice, prized for its rarity and used in various luxury applications.",
			},
			buyRangeLow:  6000,
			buyRangeHigh: 12000,
			sellDelta:    1000,
		},
		{
			Resource: Resource{
				Name:        "Copper",
				Description: "Copper is a reddish-brown metal known for its conductivity and malleability.",
			},
			buyRangeLow:  20,
			buyRangeHigh: 50,
			sellDelta:    10,
		},
		{
			Resource: Resource{
				Name:        "Silver",
				Description: "Silver is a lustrous, precious metal admired for its conductivity and versatile use in jewelry and industry.",
			},
			buyRangeLow:  250,
			buyRangeHigh: 750,
			sellDelta:    100,
		},
		{
			Resource: Resource{
				Name:        "Gold",
				Description: "Gold is a prized, yellow metal celebrated for its rarity, beauty, and historical significance in various cultural and economic contexts.",
			},
			buyRangeLow:  5000,
			buyRangeHigh: 10000,
			sellDelta:    2000,
		},
		{
			Resource: Resource{
				Name:        "Platinum",
				Description: "Platinum is a dense, silvery-white metal valued for its durability and use in jewelry and catalytic converters.",
			},
			buyRangeLow:  7500,
			buyRangeHigh: 15000,
			sellDelta:    2000,
		},
	}
}
