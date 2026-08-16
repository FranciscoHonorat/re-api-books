package valueobjects

import errD "movies-service/internal/core/domain/err-d"

type MovieTitle struct {
	Title string
}

func NewMovieTitle(title string) (*MovieTitle, error) {
	if title == "" {
		return nil, errD.ErrTitleNotValid
	}
	return &MovieTitle{Title: title}, nil
}

func (m *MovieTitle) Equals(other *MovieTitle) bool {
	if other == nil {
		return false
	}
	return m.Title == other.Title
}

func (m *MovieTitle) ZeroValue() bool {
	return m.Title == ""
}
