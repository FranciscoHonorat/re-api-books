package valueobjects_test

import (
	errD "movies-service/internal/core/domain/err-d"
	"movies-service/internal/core/domain/valueobjects"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMovieYear(t *testing.T) {
	validateMovieYear := func(year string, expectedError bool) {
		t.Helper()
		_, err := valueobjects.NewMovieYear(year)
		if expectedError {
			assert.Error(t, err)
			assert.Equal(t, errD.ErrYearNotValid, err)
		} else {
			assert.NoError(t, err)
		}
	}
	t.Run("Valid Year", func(t *testing.T) {
		validateMovieYear("Inception", false)
	})
	t.Run("Invalid Year", func(t *testing.T) {
		validateMovieYear("", true)
	})
}
