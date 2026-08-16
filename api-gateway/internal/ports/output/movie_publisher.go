package output

import (
	"context"

	"github.com/FranciscoHonorat/movies/shared"
)

type MoviePublisher interface {
	Publish(ctx context.Context, movie shared.MoviePublisherMessage) error
}
