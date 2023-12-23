package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/0z0sk0/intersvyaz-go-test/server"
	trackhandler "github.com/0z0sk0/intersvyaz-go-test/track/http/delivery"
	trackrepo "github.com/0z0sk0/intersvyaz-go-test/track/repository"
	trackusecase "github.com/0z0sk0/intersvyaz-go-test/track/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func SetUpRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = io.Discard
	router := gin.Default()

	return router
}

type TrackReq struct {
	UUID string `json:"id"`
	Page string `json:"page"`
}

var tmpUUID = uuid.New().String()

/*
unit тестирование на создание метрики, ничего сложного
*/
func TestCreateTrack(t *testing.T) {
	conn := server.LoadDB()

	r := SetUpRouter()
	trackRepository := trackrepo.NewTrackRepository(conn)
	trackUseCase := trackusecase.NewTrackUseCase(trackRepository)

	trackhandler.RegisterEndpoints(r, trackUseCase)

	jsonValue, _ := json.Marshal(TrackReq{UUID: tmpUUID, Page: "testPage"})

	req, _ := http.NewRequest("POST", "/track", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

type TrackRow struct {
	ID   int32     `db:"id"`
	UUID string    `db:"uuid"`
	Page string    `db:"page"`
	Time time.Time `db:"time"`
}

func TestCheckRow(t *testing.T) {
	conn := server.LoadDB()

	query := `SELECT * FROM track WHERE uuid = @uuid`
	args := pgx.NamedArgs{
		"uuid": tmpUUID,
	}
	row, _ := conn.Query(context.Background(), query, args)
	tracks, _ := pgx.CollectRows(row, pgx.RowToStructByName[TrackRow])
	for _, v := range tracks {
		assert.Equal(t, v.UUID, tmpUUID)
	}
}

/*
Псевдо-высоконагруженное unit тестирование сходное тесту выше
*/
func TestCreateMultipleTrack(t *testing.T) {
	conn := server.LoadDB()

	r := SetUpRouter()
	trackRepository := trackrepo.NewTrackRepository(conn)
	trackUseCase := trackusecase.NewTrackUseCase(trackRepository)

	trackhandler.RegisterEndpoints(r, trackUseCase)

	// Почему 70?
	// Из расчёта 250000 rph == ~70 rps
	for i := 0; i < 70; i++ {
		id := uuid.New()
		jsonValue, _ := json.Marshal(TrackReq{UUID: id.String(), Page: "testPage"})

		req, _ := http.NewRequest("POST", "/track", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	}
}
