package main

import (
	"github.com/labstack/echo/v4"
	"movies-service/controller"
	"movies-service/service"
	"net/http"
)

func main() {
	e := SetUp()
	e.Start(":8080")
}

func SetUp() *echo.Echo {
	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	defaultMovieService := service.DefaultMovieService{}
	movieController := controller.MovieController{
		MovieService: &defaultMovieService,
	}
	e.GET("/movies", movieController.GetMovies)

	return e
}
