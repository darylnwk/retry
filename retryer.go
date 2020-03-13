package retry

import (
	"time"
)

var (
	defaultAttempts = uint(1)
	defaultBackoff  = NoBackoff
	defaultJitter   = NoJitter
)

// Retryer defines a retryer
type Retryer struct {
	// Attempts defines the number of retry attempts.
	// Default 1.
	Attempts uint

	// Backoff defines a backoff function that returns `time.Duration`.
	// This is applied on subsequent attempts.
	// Default no backoff.
	Backoff func(n uint, delay time.Duration) time.Duration

	// Jitter defines a jitter function that returns `time.Duration`.
	// Default no jitter.
	Jitter func(backoff time.Duration) time.Duration

	// Delay defines duration to delay.
	// Default 0.
	Delay time.Duration
}

// Do executes `fn` and returns success if `fn` executed succesfully and any errors that occurred
func (retryer Retryer) Do(fn func() error) (bool, Errors) {
	var (
		n    uint
		errs Errors
	)

	if retryer.Attempts < 1 {
		retryer.Attempts = defaultAttempts
	}

	if retryer.Backoff == nil {
		retryer.Backoff = defaultBackoff
	}

	if retryer.Jitter == nil {
		retryer.Jitter = defaultJitter
	}

	for n < retryer.Attempts {
		// Backoff and jitter after first attempt only
		if n > 0 {
			time.Sleep(retryer.Jitter(retryer.Backoff(n, retryer.Delay)))
		}

		err := fn()
		if err == nil {
			return true, errs
		}

		errs = append(errs, err)

		n++
	}

	return false, errs
}
