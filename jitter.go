package retry

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NoJitter defines no jitter strategy
func NoJitter(backoff time.Duration) time.Duration {
	return backoff
}

// Jitter defines a full jitter strategy
func Jitter(backoff time.Duration) time.Duration {
	if int64(backoff) == 0 {
		return 0
	}

	return time.Duration(rand.Int63n(int64(backoff)))
}
