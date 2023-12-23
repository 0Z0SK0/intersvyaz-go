package tests

import (
	"bytes"
	"github.com/0z0sk0/intersvyaz-go-test/server"
	trackhandler "github.com/0z0sk0/intersvyaz-go-test/track/http/delivery"
	trackrepo "github.com/0z0sk0/intersvyaz-go-test/track/repository"
	trackusecase "github.com/0z0sk0/intersvyaz-go-test/track/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	// google uuid
)

func SetUpRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	router := gin.Default()

	return router
}

/*
unit тестирование на создание метрики, ничего сложного
*/
func TestCreateTrackPostReturnsStatusOk(t *testing.T) {
	conn := server.LoadDB()

	r := SetUpRouter()
	trackRepository := trackrepo.NewTrackRepository(conn)
	trackUseCase := trackusecase.NewTrackUseCase(trackRepository)

	trackhandler.RegisterEndpoints(r, trackUseCase)

	var jsonValue = []byte(`{"id":"60b496c1", "page":"demo"}`)

	req, _ := http.NewRequest("POST", "/track", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

/*
type TrackReq struct {
	id uuid
	page string
}
*/

/*
Псевдо-высоконагруженное unit тестирование сходное тесту выше
*/
func TestHighLoadCreateTrackPostReturnsStatusOk(t *testing.T) {
	conn := server.LoadDB()

	r := SetUpRouter()
	trackRepository := trackrepo.NewTrackRepository(conn)
	trackUseCase := trackusecase.NewTrackUseCase(trackRepository)

	trackhandler.RegisterEndpoints(r, trackUseCase)

	// Чтобы избежать кеширования, лучше заменить готовую строку на генерированную
	var jsonValue = []byte(`{"id":"60b496c1", "page":"demo"}`)

        /*
	trackReq := TrackReq{
            id: uuid.String(),
	    page: "testPage",
        }
	var jsonValue = json.Marshal(trackReq)
	*/

	// Почему 70?
	// Из расчёта 250000 rph == ~70 rps
	for i := 0; i < 70; i++ {
		req, _ := http.NewRequest("POST", "/track", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	}
}
