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

// Commodity resources.
var (
// TODO: Commodity resources
)

// Equipment resources.
var (
// TODO: Equipment resources
)

// Land resources.
var (
// TODO: Land resources
)

// Employee resources.
var (
	RESOURCE_WORKER = &Resource{
		Name:        "Worker",
		Description: "",
		Type:        RESOURCE_TYPE_EMPLOYEE,
	}

	RESOURCE_SURVEYOR = &Resource{
		Name:        "Surveyor",
		Description: "",
		Type:        RESOURCE_TYPE_EMPLOYEE,
	}

	RESOURCE_SPECIALIST = &Resource{
		Name:        "Specialist",
		Description: "",
		Type:        RESOURCE_TYPE_EMPLOYEE,
	}

	RESOURCE_MINING_ENGINEER = &Resource{
		Name:        "Mining Engineer",
		Description: "",
		Type:        RESOURCE_TYPE_EMPLOYEE,
	}
)
