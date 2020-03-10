package retry

import "time"

// NoBackoff defines no backoff strategy
func NoBackoff() time.Duration {
	return time.Duration(0)
}
