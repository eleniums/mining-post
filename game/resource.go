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
}
