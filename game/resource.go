package game

type ResourceType string

const (
	RESOURCE_TYPE_COMMODITY ResourceType = "Commodity"
	RESOURCE_TYPE_EQUIPMENT ResourceType = "Equipment"
	RESOURCE_TYPE_LAND      ResourceType = "Land"
	RESOURCE_TYPE_EMPLOYEE  ResourceType = "Employee"
)

type Resource struct {
	Name        string
	Description string
	Type        ResourceType
}
