package usecase

import (
	"context"
	"github.com/0z0sk0/intersvyaz-go-test/models"
	"github.com/0z0sk0/intersvyaz-go-test/track/repository"
)

type TrackUseCase struct {
	trackRepo repository.TrackRepository
}

func NewTrackUseCase(trackRepo *repository.TrackRepository) *TrackUseCase {
	return &TrackUseCase{
		trackRepo: *trackRepo,
	}
}

func (track TrackUseCase) CreateTrack(ctx context.Context, uuid, page string) error {
	tr := &models.Track{
		UUID: uuid,
		Page: page,
	}

	return track.trackRepo.CreateTrack(ctx, tr)
}
