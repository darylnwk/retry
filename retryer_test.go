package retry_test

import (
	"errors"
	"testing"

	"github.com/darylnwk/retry"
	"github.com/stretchr/testify/assert"
)

func TestRetry_Do_Success(t *testing.T) {
	retryer := retry.Retryer{}
	ok, err := retryer.Do(func() error {
		return nil
	})

	assert := assert.New(t)
	assert.True(ok)
	assert.Equal(0, len(err))
}

func TestRetry_Do_FailWithErrors(t *testing.T) {
	retryer := retry.Retryer{
		Attempts: 2,
		Backoff:  retry.NoBackoff,
	}
	ok, err := retryer.Do(func() error {
		return errors.New("something bad happened")
	})

	assert := assert.New(t)
	assert.False(ok)
	assert.Equal(2, len(err))
	assert.Equal("something bad happened; something bad happened", err.Error())
}
