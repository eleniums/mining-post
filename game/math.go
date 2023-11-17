package game

import (
	"log/slog"
	"math/rand"
	"strconv"
)

// Returns a random value in the range [low, high] inclusive.
func randInt64(low int64, high int64) int64 {
	return rand.Int63n(high-low+1) + low
}

// Returns a random value in the range [low, high] inclusive (includes 2 decimal places).
func randFloat64(low float64, high float64) float64 {
	hundredsPlace := randInt64(int64(low*100.0), int64(high*100.0))
	return float64(hundredsPlace) / 100.0
}

// Rounds a value to the given precision.
func roundFloat64(val float64, precision int) float64 {
	s := strconv.FormatFloat(val, 'f', precision, 64)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		slog.Error("error parsing float for rounding operation", ErrAttr(err))
		return 0
	}
	return f
}
