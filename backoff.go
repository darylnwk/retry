package retry

import "time"

// NoBackoff defines no backoff strategy
func NoBackoff(_ uint, _ time.Duration) time.Duration {
	return time.Duration(0)
}

// ConstantBackoff defines constant backoff strategy
func ConstantBackoff(_ uint, delay time.Duration) time.Duration {
	return delay
}

// ExponentialBackoff defines exponential backoff strategy
func ExponentialBackoff(n uint, delay time.Duration) time.Duration {
	return delay * (1 << n)
}
