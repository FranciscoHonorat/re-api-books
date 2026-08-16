package input

import (
	"context"
	"movies-service/internal/core/domain/entity"
	"movies-service/internal/core/port/output"
)

type MovieService interface {
	GetMovieByID(id int) (*entity.MovieEntity, error)
	ListMovies(filters output.Listfilters, pagination output.Pagination, sorting output.Sorting) ([]*entity.MovieEntity, error)
	CountMovies(filters output.Listfilters) (int, error)
	CreateMovie(ctx context.Context, movie *entity.MovieEntity) error
	DeleteMovie(id int) error
}
