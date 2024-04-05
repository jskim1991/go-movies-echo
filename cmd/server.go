package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"movie-service/internal/controller"
	"movie-service/pkg/db"
	"movie-service/pkg/repository"
	"movie-service/pkg/service"
	"net/http"
)

func main() {
	e := configureEcho()
	e.Start(":8080")
}

func configureEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	dsn := "postgres://postgres:@localhost:5432/movies"
	postgres := db.NewPostgresConnection(dsn)
	defaultMovieRepository := repository.NewDefaultMovieRepository(postgres)
	defaultMovieService := service.NewDefaultMovieService(defaultMovieRepository)
	movieController := controller.NewMovieController(defaultMovieService)
	e.GET("/movies", movieController.GetMovies)

	return e
}
