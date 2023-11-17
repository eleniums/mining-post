package game

func GetInitialStock() []Listing {
	stock := []Listing{
		{
			Resource: Resource{
				Name:        "Granite",
				Description: "Rough hewn igneous rock that is rather durable.",
			},
			Quantity:  1000,
			BuyPrice:  10.00,
			SellPrice: 8.00,
		},
		{
			Resource: Resource{
				Name:        "Limestone",
				Description: "Soft, chalky, sedimentary rock that is easily crushed for use in concrete",
			},
			Quantity:  1000,
			BuyPrice:  4.00,
			SellPrice: 8.00,
		},
		{
			Resource: Resource{
				Name:        "Marble",
				Description: "Metamorphic rock that polishes well and is often used for decorative purposes.",
			},
			Quantity:  1000,
			BuyPrice:  20.00,
			SellPrice: 8.00,
		},
		{
			Resource: Resource{
				Name:        "Sandstone",
				Description: "Sedimentary rock unsuitable for buildings that is primarily used for concrete.",
			},
			Quantity:  1000,
			BuyPrice:  8.00,
			SellPrice: 8.00,
		},
	}

	return stock
}
