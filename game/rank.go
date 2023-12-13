package game

type Rank struct {
	Name   string
	Salary float64

	eligibleForPromotion func(player *Player) bool
}

// TODO: make sure salaries and promotions make sense
var ranks = []Rank{
	{
		Name:   "Amateur Miner L1",
		Salary: 10,
		eligibleForPromotion: func(player *Player) bool {
			return player.Money >= 250
		},
	},
	{
		Name:   "Amateur Miner L2",
		Salary: 50,
		eligibleForPromotion: func(player *Player) bool {
			return player.Money >= 500
		},
	},
	{
		Name:   "Amateur Miner L3",
		Salary: 75,
		eligibleForPromotion: func(player *Player) bool {
			return player.Money >= 1_000
		},
	},
	{
		Name:   "Apprentice Miner L1",
		Salary: 100,
		eligibleForPromotion: func(player *Player) bool {
			return player.Money >= 10_000
		},
	},
	{
		Name:   "Apprentice Miner L2",
		Salary: 250,
		eligibleForPromotion: func(player *Player) bool {
			return player.Money >= 25_000
		},
	},
	{
		Name:   "Apprentice Miner L3",
		Salary: 500,
		eligibleForPromotion: func(player *Player) bool {
			return player.Money >= 50_000
		},
	},
	{
		Name:   "Miner L1",
		Salary: 1_000,
		eligibleForPromotion: func(player *Player) bool {
			return player.Money >= 100_000
		},
	},
	{
		Name:   "Miner L2",
		Salary: 5_000,
		eligibleForPromotion: func(player *Player) bool {
			return player.Money >= 250_000
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
			return player.Money >= 100_000_000
		},
	},
	{
		Name:   "Distinguished Miner",
		Salary: 750_000,
		eligibleForPromotion: func(player *Player) bool {
			return player.Money >= 1_000_000_000
		},
	},
	{
		Name:   "Legendary Miner",
		Salary: 1_000_000,
		eligibleForPromotion: func(player *Player) bool {
			return player.Money >= 1_000_000_000_000
		},
	},
	{
		Name:   "Mythic Miner",
		Salary: 10_000_000,
		eligibleForPromotion: func(player *Player) bool {
			return false
		},
	},
}
