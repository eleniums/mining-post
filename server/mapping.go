package server

import (
	"github.com/eleniums/mining-post/mem"
	"github.com/eleniums/mining-post/models"
)

func mapItemToDBItem(in *models.Item) *mem.Item {
	return &mem.Item{
		Category: in.Category,
		Data:     in.Data,
		Tags:     in.Tags,
		Metadata: in.Metadata,
	}
}

func mapDBItemToItem(id string, in *mem.Item) *models.Item {
	return &models.Item{
		ID:       id,
		Category: in.Category,
		Data:     in.Data,
		Tags:     in.Tags,
		Metadata: in.Metadata,
	}
}
