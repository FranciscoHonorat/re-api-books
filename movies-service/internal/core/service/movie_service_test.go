package service_test

import (
	"context"
	"movies-service/internal/core/domain/entity"
	"movies-service/internal/core/port/output"
	"movies-service/internal/core/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMovieRepository struct {
	mock.Mock
}

func (m *MockMovieRepository) GetMovieByID(ctx context.Context, id int) (*entity.MovieEntity, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.MovieEntity), args.Error(1)
}

func (m *MockMovieRepository) ListMovies(ctx context.Context, filters output.Listfilters, pagination output.Pagination, sorting output.Sorting) ([]entity.MovieEntity, error) {
	args := m.Called(ctx, filters, pagination, sorting)
	return args.Get(0).([]entity.MovieEntity), args.Error(1)
}

func (m *MockMovieRepository) CountMovies(ctx context.Context, filters output.Listfilters) (int, error) {
	args := m.Called(ctx, filters)
	return args.Int(0), args.Error(1)
}

func (m *MockMovieRepository) CreateMovie(ctx context.Context, movie *entity.MovieEntity) (*entity.MovieEntity, error) {
	args := m.Called(ctx, movie)
	return args.Get(0).(*entity.MovieEntity), args.Error(1)
}

func (m *MockMovieRepository) DeleteMovie(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestMovieService(t *testing.T) {
	mockRepo := new(MockMovieRepository)
	movieService := service.NewMovieService(mockRepo)

	t.Run("GetMovieByID", func(t *testing.T) {
		movie1, err := entity.NewMovieEntity(1, "Movie 1", "2024")
		assert.NoError(t, err)
		movie2, err := entity.NewMovieEntity(2, "Movie 2", "2023")
		assert.NoError(t, err)
		mockMovies := []entity.MovieEntity{*movie1, *movie2}
		mockRepo.On("ListMovies", mock.Anything, output.Listfilters{}, output.Pagination{Page: 1, Limit: 10}, output.Sorting{}).Return(mockMovies, nil)
		mockRepo.On("CountMovies", mock.Anything, output.Listfilters{}).Return(1, nil)

		movies, total, err := movieService.ListMovies(context.Background(), output.Pagination{Page: 1, Limit: 10}, output.Listfilters{}, output.Sorting{})
		assert.NoError(t, err)
		assert.Equal(t, mockMovies, movies)
		assert.Equal(t, 1, total)
		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateMovie", func(t *testing.T) {
		newMovie, err := entity.NewMovieEntity(0, "New Movie", "2024")
		assert.NoError(t, err)
		mockRepo.On("CreateMovie", mock.Anything, newMovie).Return(newMovie, nil)

		createdMovie, err := movieService.CreateMovie(context.Background(), newMovie)
		assert.NoError(t, err)
		assert.Equal(t, newMovie, createdMovie)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteMovie", func(t *testing.T) {
		mockRepo.On("DeleteMovie", mock.Anything, 1).Return(nil)

		err := movieService.DeleteMovie(context.Background(), 1)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}
