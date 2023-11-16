package models

type Item struct {
	ID       string            `json:"id,omitempty"`
	Category int32             `json:"category,omitempty"`
	Data     string            `json:"data,omitempty"`
	Tags     []string          `json:"tags,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

type Items struct {
	Items []*Item `json:"items,omitempty"`
}
