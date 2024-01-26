package game

import (
	"fmt"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_Integration_CalculateLoot_Success(t *testing.T) {
	loot := LootTable{
		{
			Name:      "Loot1",
			Weight:    25,
			CountLow:  1,
			CountHigh: 3,
		},
		{
			Name:      "Loot2",
			Weight:    65,
			CountLow:  1,
			CountHigh: 3,
		},
		{
			Name:      "Loot3",
			Weight:    10,
			CountLow:  1,
			CountHigh: 3,
		},
	}

	nameTally := map[string]int{}
	quantityTally := map[int64]int{}

	iterations := 10_000
	for i := 0; i < iterations; i++ {
		name, quantity := loot.CalculateLoot()

		nameTally[name]++
		quantityTally[quantity]++
	}

	fmt.Println(nameTally)
	fmt.Println(quantityTally)

	assertPercentWithinRange(t, nameTally[loot[0].Name], iterations, float64(loot[0].Weight))
	assertPercentWithinRange(t, nameTally[loot[1].Name], iterations, float64(loot[1].Weight))
	assertPercentWithinRange(t, nameTally[loot[2].Name], iterations, float64(loot[2].Weight))
}

func assertPercentWithinRange(t *testing.T, count int, iterations int, targetPercent float64) {
	variability := 2.0
	percent := float64(count) / float64(iterations) * 100.0
	assert.True(t, percent > targetPercent-variability && percent < targetPercent+variability)
}
