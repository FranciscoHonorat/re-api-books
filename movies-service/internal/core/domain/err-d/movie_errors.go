package errD

import "errors"

var (
	ErrMovieNotFound = errors.New("Movie not found")
	ErrIDNotValid    = errors.New("id not valid")
	ErrTitleNotValid = errors.New("title not valid")
	ErrYearNotValid  = errors.New("year not valid")
)
