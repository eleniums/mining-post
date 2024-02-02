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

	// Loot table will be processed every world update. Used to generate
	// resources from mines.
	Loot LootTable

	buyRangeLow  float64 // lowest possible buy price for resource
	buyRangeHigh float64 // highest possible buy price for resource
	sellDelta    float64 // highest potential difference of selling price. Selling price is always lower than buying price
	netWorth     float64 // this is calculated based on the average of the buy low and buy high and includes prereqs
}

// Prerequisite for purchasing a resource. Includes name of resource and amount to subtract from player inventory.
type Prerequisite struct {
	Name     string
	Quantity int64
}

// This will calculate and return net worth for a resource.
func (r *Resource) CalculateNetWorth() float64 {
	// only calculate net worth if it has not been calculated already
	if r.netWorth > 0 {
		return r.netWorth
	}

	// calculate total worth of prereqs
	var prereqTotal float64 = 0
	for _, prereq := range r.Prerequisites {
		resource := findResource(prereq.Name)
		if resource != nil {
			prereqTotal += resource.CalculateNetWorth() * float64(prereq.Quantity)
		}
	}

	// average buy low/high and add prereq total
	r.netWorth = (r.buyRangeLow+r.buyRangeHigh)/2.0 + prereqTotal

	return r.netWorth
}
