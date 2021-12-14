package retry_test

import (
	stderrors "errors"
	"testing"

	"github.com/IQ-tech/go-errors"
	"github.com/IQ-tech/go-retry"
	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	t.Parallel()

	t.Run("returns error if attempts are exceeded", func(t *testing.T) {
		t.Parallel()

		err := stderrors.New("message")

		function := func() (interface{}, error) {
			return nil, err
		}

		_, retryErr := retry.Func(retry.Options{Attempts: 3, InitialTimeBetweenRetries: 1}, function)

		assert.Equal(t, errors.GetOriginalError(retryErr).Error(), err.Error())
	})

	t.Run("doesn't return error if function succeeds before attempts are exceeded", func(t *testing.T) {
		t.Parallel()

		err := stderrors.New("message")

		attempts := 0

		function := func() (interface{}, error) {
			attempts++

			if attempts == 2 {
				return 1, nil
			}

			return nil, err
		}

		result, err := retry.Func(retry.Options{Attempts: 3, InitialTimeBetweenRetries: 1}, function)

		assert.NoError(t, err)

		assert.Equal(t, 1, result)
	})

	t.Run("doesn't return error if function succeeds in the first try", func(t *testing.T) {
		t.Parallel()

		function := func() (interface{}, error) {
			return 2, nil
		}

		result, err := retry.Func(retry.Options{Attempts: 5, InitialTimeBetweenRetries: 5}, function)

		assert.Nil(t, err)

		assert.Equal(t, 2, result)
	})
}
