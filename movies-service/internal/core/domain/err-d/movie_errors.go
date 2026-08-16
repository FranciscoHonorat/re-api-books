package errD

import "errors"

var (
	ErrMovieNotFound    = errors.New("movie not found")
	ErrIDNotValid       = errors.New("id not valid")
	ErrTitleNotValid    = errors.New("title not valid")
	ErrYearNotValid     = errors.New("year not valid")
	ErrInvalidMovieData = errors.New("invalid movie data")
)
