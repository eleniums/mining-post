package game

// Get default list of all stock and prices.
func getInitialStock() []*Listing {
	return []*Listing{
		{
			Resource: Resource{
				Name:        "Limestone",
				Description: "Soft, chalky, sedimentary rock that is easily crushed for use in concrete",
			},
			buyRangeLow:  5,
			buyRangeHigh: 20,
			sellDelta:    4,
		},
		{
			Resource: Resource{
				Name:        "Sandstone",
				Description: "Sedimentary rock unsuitable for buildings that is primarily used for concrete.",
			},
			buyRangeLow:  8,
			buyRangeHigh: 30,
			sellDelta:    7,
		},
		{
			Resource: Resource{
				Name:        "Granite",
				Description: "Rough hewn igneous rock that is rather durable.",
			},
			buyRangeLow:  5,
			buyRangeHigh: 50,
			sellDelta:    4,
		},
		{
			Resource: Resource{
				Name:        "Marble",
				Description: "Metamorphic rock that polishes well and is often used for decorative purposes.",
			},
			buyRangeLow:  10,
			buyRangeHigh: 100,
			sellDelta:    9,
		},
	}
}
