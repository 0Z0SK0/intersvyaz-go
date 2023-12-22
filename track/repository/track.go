package repository

import (
	"context"
	"github.com/0z0sk0/intersvyaz-go-test/models"
	"github.com/jackc/pgx/v5"
)

type TrackRepository struct {
	db *pgx.Conn
}

func NewTrackRepository(db *pgx.Conn) *TrackRepository {
	return &TrackRepository{
		db: db,
	}
}

func (repo TrackRepository) CreateTrack(ctx context.Context, track *models.Track) error {
	query := `INSERT INTO track (uuid, page, time) VALUES (@uuid, @page, NOW())`
	args := pgx.NamedArgs{
		"uuid": track.UUID,
		"page": track.Page,
	}

	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return err
	}

	return nil
}
