package controller

import (
	"github.com/labstack/echo/v4"
	"movies-service/service"
	"net/http"
)

type MovieController struct {
	MovieService service.MovieService
}

func (m *MovieController) GetMovies(c echo.Context) error {
	movies, _ := m.MovieService.FetchMovies()

	return c.JSON(http.StatusOK, movies)
}
