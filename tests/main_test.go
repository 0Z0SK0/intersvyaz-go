package tests

import (
	"bytes"
	"encoding/json"
	"github.com/0z0sk0/intersvyaz-go-test/server"
	trackhandler "github.com/0z0sk0/intersvyaz-go-test/track/http/delivery"
	trackrepo "github.com/0z0sk0/intersvyaz-go-test/track/repository"
	trackusecase "github.com/0z0sk0/intersvyaz-go-test/track/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
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

/*
unit тестирование на создание метрики, ничего сложного
*/
func TestCreateTrack(t *testing.T) {
	conn := server.LoadDB()

	r := SetUpRouter()
	trackRepository := trackrepo.NewTrackRepository(conn)
	trackUseCase := trackusecase.NewTrackUseCase(trackRepository)

	trackhandler.RegisterEndpoints(r, trackUseCase)

	id := uuid.New()
	jsonValue, _ := json.Marshal(TrackReq{UUID: id.String(), Page: "testPage"})

	var jsonDecoded, _ = json.MarshalIndent(TrackReq{UUID: id.String(), Page: "testPage"}, "", "")
	log.Printf("%s", string(jsonDecoded))

	req, _ := http.NewRequest("POST", "/track", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
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
