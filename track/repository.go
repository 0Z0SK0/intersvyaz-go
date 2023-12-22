package track

import (
	"context"
	"github.com/0z0sk0/intersvyaz-go-test/models"
)

type Repository interface {
	CreateTrack(ctx context.Context, track *models.Track) error
}
