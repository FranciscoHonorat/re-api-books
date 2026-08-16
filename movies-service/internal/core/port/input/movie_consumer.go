package input

import (
	"context"
	"movies-service/internal/core/domain/entity"
)

type MovieConsumer interface {
	ConsumeMovie(ctx context.Context, handler func(movie *entity.MovieEntity) error) error
}
