package http

import (
	"github.com/0z0sk0/simple-metrika-app/track"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Handler struct {
	useCase track.UseCase
}

func CreateHandler(useCase track.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

// Пост трека
// При масштабировании, увеличить на требуемое кол-во параметров
type createTrackInput struct {
	UUID string `json:"id"`
	Page string `json:"page"`
}

func (h *Handler) Create(ctx *gin.Context) {
	trackInput := new(createTrackInput)
	if err := ctx.BindJSON(trackInput); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.useCase.CreateTrack(ctx.Request.Context(), trackInput.UUID, trackInput.Page); err != nil {
		log.Printf("%s", err.Error())
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}
