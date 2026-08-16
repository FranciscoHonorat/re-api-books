package service

import (
	"context"
	"fmt"
	"movies-service/internal/core/domain/entity"
	errD "movies-service/internal/core/domain/err-d"
	"movies-service/internal/core/port/output"
)

type MovieService struct {
	repo output.MovieRepository
}

func NewMovieService(repo output.MovieRepository) *MovieService {
	return &MovieService{
		repo: repo,
	}
}

func (m *MovieService) GetMovie(ctx context.Context, id int) (*entity.MovieEntity, error) {
	if id <= 0 {
		return nil, errD.ErrIDNotValid
	}

	movie, err := m.repo.GetMovieByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (m *MovieService) ListMovies(ctx context.Context, page output.Pagination, list output.Listfilters, sorting output.Sorting) ([]entity.MovieEntity, int, error) {
	movies, err := m.repo.ListMovies(ctx, list, page, sorting)
	if err != nil {
		return nil, 0, err
	}

	total, err := m.repo.CountMovies(ctx, list)
	if err != nil {
		return nil, 0, err
	}

	return movies, total, nil
}

func (m *MovieService) CreateMovie(ctx context.Context, movie *entity.MovieEntity) (*entity.MovieEntity, error) {
	if movie == nil {
		return nil, fmt.Errorf("Invalid Movie")
	}

	if err := movie.Validate(); err != nil {
		return nil, fmt.Errorf("Invalid Movie")
	}

	createMovie, err := m.repo.CreateMovie(ctx, movie)
	if err != nil {
		return nil, fmt.Errorf("Movie no create")
	}

	return createMovie, nil
}

func (m *MovieService) DeleteMovie(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("Invalid ID")
	}

	if err := m.repo.DeleteMovie(ctx, id); err != nil {
		return err
	}

	return nil
}
