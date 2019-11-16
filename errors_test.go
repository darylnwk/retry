package retry_test

import (
	"errors"
	"testing"

	"github.com/darylnwk/retry"
	"github.com/stretchr/testify/assert"
)

func TestErrors_Error(t *testing.T) {
	errs := retry.Errors([]error{errors.New("something bad happened"), errors.New("this is bad")})
	assert := assert.New(t)
	assert.Equal("something bad happened; this is bad", errs.Error())
}
