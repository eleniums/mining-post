package game

import "github.com/eleniums/mining-post/models"

func (m *Manager) GetMarketStock() models.Market {
	return m.market
}
