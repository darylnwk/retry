package retry

// Retryer defines a retryer
type Retryer struct {
	// Attempts defines the number of retry attempts
	Attempts uint
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

	for n < retryer.Attempts {
		err := fn()
		if err == nil {
			return true, errs
		}

		errs = append(errs, err)

		n++
	}

	return false, errs
}
