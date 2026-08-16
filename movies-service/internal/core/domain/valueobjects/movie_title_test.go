package valueobjects_test

import (
	errD "movies-service/internal/core/domain/err-d"
	"movies-service/internal/core/domain/valueobjects"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMovieTitle(t *testing.T) {
	validateMovieTitle := func(title string, expectedError bool) {
		t.Helper()
		_, err := valueobjects.NewMovieTitle(title)
		if expectedError {
			assert.Error(t, err)
			assert.Equal(t, errD.ErrTitleNotValid, err)
		} else {
			assert.NoError(t, err)
		}
	}
	t.Run("Valid Title", func(t *testing.T) {
		validateMovieTitle("Inception", false)
	})
	t.Run("Invalid Title", func(t *testing.T) {
		validateMovieTitle("", true)
	})
}
