package client

import (
	"math/rand"
	"time"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/backoff"
	"github.com/Rican7/retry/jitter"
	"github.com/Rican7/retry/strategy"
)

const (
	MaxRetries      = 3
	BackoffDuration = 200 * time.Millisecond
)

var generator = rand.New(rand.NewSource(time.Now().UnixNano()))

func Retry(action func(attempt uint) error) {
	// TODO: integrate this into client somehow
	retry.Retry(
		action,
		strategy.Limit(MaxRetries+1), // strategy.Limit is the total number of attempts, so original attempt + max retries
		strategy.BackoffWithJitter(
			backoff.BinaryExponential(BackoffDuration),
			jitter.Deviation(generator, 0.5),
		),
	)
}
