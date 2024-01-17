package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/0z0sk0/simple-metrika-app/track"
	trackhttp "github.com/0z0sk0/simple-metrika-app/track/http/delivery"
	trackrepo "github.com/0z0sk0/simple-metrika-app/track/repository"
	trackusecase "github.com/0z0sk0/simple-metrika-app/track/usecase"
)

type App struct {
	http *http.Server
	DB   *pgx.Conn

	// usecases must be placed there
	trackUseCase track.UseCase
}

func CreateApp() *App {
	db := LoadDB()

	// create Repo
	trackRepository := trackrepo.NewTrackRepository(db)

	return &App{
		DB:           db,
		trackUseCase: trackusecase.NewTrackUseCase(trackRepository),
	}
}

func (app *App) Start() error {
	// to prod
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	/*
		В идеале, здесь должен быть расположен router.Group
		с middleware в виде /api

		middleware := mw.NewMiddleware(app.middlewareUsecase)
		api := router.Group("/api", middleware)
	*/

	trackhttp.RegisterEndpoints(router, app.trackUseCase)

	app.http = &http.Server{
		Addr:           ":" + os.Getenv("APP_PORT"),
		Handler:        router,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := app.http.ListenAndServe(); err != nil {
			log.Fatalf("Unable to listen: %s", err.Error())
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return app.http.Shutdown(ctx)
}

func LoadDB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}
