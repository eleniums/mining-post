package game

type Rank struct {
	Name   string
	Salary float64

	eligibleForPromotion func(player *Player) bool
}

var ranks = []Rank{
	{
		Name:   "Miner L1",
		Salary: 1_000,
		eligibleForPromotion: func(player *Player) bool {
			return player.Money >= 50_000
		},
	},
	{
		Name:   "Miner L2",
		Salary: 5_000,
		eligibleForPromotion: func(player *Player) bool {
			return player.Money >= 200_000
		},
	},
	{
		Name:   "Miner L3",
		Salary: 10_000,
		eligibleForPromotion: func(player *Player) bool {
			return player.Money >= 500_000
		},
	},
	{
		Name:   "Senior Miner",
		Salary: 30_000,
		eligibleForPromotion: func(player *Player) bool {
			return player.Money >= 1_000_000
		},
	},
	{
		Name:   "Staff Miner",
		Salary: 100_000,
		eligibleForPromotion: func(player *Player) bool {
			return player.Money >= 10_000_000
		},
	},
	{
		Name:   "Principal Miner",
		Salary: 250_000,
		eligibleForPromotion: func(player *Player) bool {
			return player.Money >= 1_000_000_000
		},
	},
	{
		Name:   "Distinguished Miner",
		Salary: 750_000,
		eligibleForPromotion: func(player *Player) bool {
			return false
		},
	},
}
