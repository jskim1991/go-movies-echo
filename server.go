package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"movie-service/controller"
	"movie-service/repository"
	"movie-service/service"
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
	postgres := repository.NewPostgresConnection(dsn)
	defaultMovieRepository := repository.NewDefaultMovieRepository(postgres)
	defaultMovieService := service.NewDefaultMovieService(defaultMovieRepository)
	movieController := controller.NewMovieController(defaultMovieService)
	e.GET("/movies", movieController.GetMovies)

	return e
}
