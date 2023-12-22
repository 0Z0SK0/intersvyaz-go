package track

import (
	"context"
)

type UseCase interface {
	CreateTrack(ctx context.Context, uuid, page string) error
}
