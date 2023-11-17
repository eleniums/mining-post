package game

// Get default list of all stock and prices.
func GetInitialStock() []Listing {
	stock := []Listing{
		{
			Resource: Resource{
				Name:        "Limestone",
				Description: "Soft, chalky, sedimentary rock that is easily crushed for use in concrete",
			},
			quantityRangeLow:  10,
			quantityRangeHigh: 1000,
			buyRangeLow:       5,
			buyRangeHigh:      20,
			sellDelta:         4,
		},
		{
			Resource: Resource{
				Name:        "Sandstone",
				Description: "Sedimentary rock unsuitable for buildings that is primarily used for concrete.",
			},
			quantityRangeLow:  10,
			quantityRangeHigh: 1000,
			buyRangeLow:       8,
			buyRangeHigh:      30,
			sellDelta:         7,
		},
		{
			Resource: Resource{
				Name:        "Granite",
				Description: "Rough hewn igneous rock that is rather durable.",
			},
			quantityRangeLow:  10,
			quantityRangeHigh: 1000,
			buyRangeLow:       5,
			buyRangeHigh:      50,
			sellDelta:         4,
		},
		{
			Resource: Resource{
				Name:        "Marble",
				Description: "Metamorphic rock that polishes well and is often used for decorative purposes.",
			},
			quantityRangeLow:  10,
			quantityRangeHigh: 750,
			buyRangeLow:       10,
			buyRangeHigh:      100,
			sellDelta:         9,
		},
	}

	return stock
}
