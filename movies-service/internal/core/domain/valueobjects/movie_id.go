package valueobjects

import errD "movies-service/internal/core/domain/err-d"

type MovieID struct {
	id int
}

func NewMovieID(id int) (*MovieID, error) {
	if id <= 0 {
		return nil, errD.ErrIDNotValid
	}
	return &MovieID{id: id}, nil
}

func (m *MovieID) ID() int {
	return m.id
}

func (m *MovieID) Equals(other *MovieID) bool {
	if other == nil {
		return false
	}
	return m.id == other.id
}

func (m *MovieID) ZeroValue() bool {
	return m.id == 0
}
