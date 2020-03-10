package retry

import "time"

// Retryer defines a retryer
type Retryer struct {
	// Attempts defines the number of retry attempts
	Attempts uint

	// Backoff defines a backoff function that returns `time.Duration`
	// This is applied on subsequent attempts.
	Backoff func() time.Duration
}

// Do executes `fn` and returns success if `fn` executed succesfully and any errors that occurred
func (retryer Retryer) Do(fn func() error) (bool, Errors) {
	var (
		n    uint
		errs Errors
	)

	if retryer.Attempts < 1 {
		retryer.Attempts = 1
	}

	if retryer.Backoff == nil {
		retryer.Backoff = NoBackoff
	}

	for n < retryer.Attempts {
		// Backoff after first attempt only
		if n > 0 {
			time.Sleep(retryer.Backoff())
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
