package valueobjects

import errD "movies-service/internal/core/domain/err-d"

type MovieYear struct {
	Year string
}

func NewMovieYear(year string) (*MovieYear, error) {
	if year == "" {
		return nil, errD.ErrYearNotValid
	}
	return &MovieYear{Year: year}, nil
}

func (m *MovieYear) Equals(other *MovieYear) bool {
	if other == nil {
		return false
	}
	return m.Year == other.Year
}

func (m *MovieYear) ZeroValue() bool {
	return m.Year == ""
}
