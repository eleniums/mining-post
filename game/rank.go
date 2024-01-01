package game

type Rank struct {
	Name   string
	Salary float64

	eligibleForPromotion func(player *Player) bool
}

// Master list of all ranks with associated salaries.
var ranks = []Rank{
	{
		Name:   "Amateur Miner L1",
		Salary: 10,
		eligibleForPromotion: func(player *Player) bool {
			return player.NetWorth >= 1000
		},
	},
	{
		Name:   "Amateur Miner L2",
		Salary: 50,
		eligibleForPromotion: func(player *Player) bool {
			return player.NetWorth >= 5000
		},
	},
	{
		Name:   "Amateur Miner L3",
		Salary: 100,
		eligibleForPromotion: func(player *Player) bool {
			return player.NetWorth >= 10_000
		},
	},
	{
		Name:   "Apprentice Miner L1",
		Salary: 500,
		eligibleForPromotion: func(player *Player) bool {
			return player.NetWorth >= 50_000
		},
	},
	{
		Name:   "Apprentice Miner L2",
		Salary: 1000,
		eligibleForPromotion: func(player *Player) bool {
			return player.NetWorth >= 100_000
		},
	},
	{
		Name:   "Apprentice Miner L3",
		Salary: 2000,
		eligibleForPromotion: func(player *Player) bool {
			return player.NetWorth >= 200_000
		},
	},
	{
		Name:   "Miner L1",
		Salary: 10_000,
		eligibleForPromotion: func(player *Player) bool {
			return player.NetWorth >= 1_000_000
		},
	},
	{
		Name:   "Miner L2",
		Salary: 20_000,
		eligibleForPromotion: func(player *Player) bool {
			return player.NetWorth >= 2_000_000
		},
	},
	{
		Name:   "Miner L3",
		Salary: 40_000,
		eligibleForPromotion: func(player *Player) bool {
			return player.NetWorth >= 4_000_000
		},
	},
	{
		Name:   "Senior Miner",
		Salary: 100_000,
		eligibleForPromotion: func(player *Player) bool {
			return player.NetWorth >= 20_000_000
		},
	},
	{
		Name:   "Staff Miner",
		Salary: 250_000,
		eligibleForPromotion: func(player *Player) bool {
			return player.NetWorth >= 50_000_000
		},
	},
	{
		Name:   "Principal Miner",
		Salary: 500_000,
		eligibleForPromotion: func(player *Player) bool {
			return player.NetWorth >= 1_000_000_000
		},
	},
	{
		Name:   "Distinguished Miner",
		Salary: 750_000,
		eligibleForPromotion: func(player *Player) bool {
			return player.NetWorth >= 1_000_000_000_000
		},
	},
	{
		Name:   "Legendary Miner",
		Salary: 1_000_000,
		eligibleForPromotion: func(player *Player) bool {
			return player.NetWorth >= 100_000_000_000_000
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
