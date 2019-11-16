package retry_test

import (
	"errors"
	"testing"

	"github.com/darylnwk/retry"
	"github.com/stretchr/testify/assert"
)

func TestRetry_Do_SuccessWithNoErrors(t *testing.T) {
	retryer := retry.Retryer{}
	ok, err := retryer.Do(func() error {
		return nil
	})

	assert := assert.New(t)
	assert.True(ok)
	assert.Equal(0, len(err))
}

func TestRetry_Do_FailWithErrors(t *testing.T) {
	retryer := retry.Retryer{}
	ok, err := retryer.Do(func() error {
		return errors.New("something bad happened")
	})

	assert := assert.New(t)
	assert.False(ok)
	assert.Equal(1, len(err))
	assert.Equal("something bad happened", err.Error())
}
