package retry

import "strings"

// Errors defines a list of `error`
type Errors []error

func (errs Errors) Error() string {
	var errStr []string
	for _, err := range errs {
		errStr = append(errStr, err.Error())
	}

	return strings.Join(errStr, "; ")
}
