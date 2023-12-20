package controller

import (
	"github.com/labstack/echo/v4"
	"movie-service/service"
	"net/http"
)

type MovieController struct {
	MovieService service.MovieService
}

func (m *MovieController) GetMovies(c echo.Context) error {
	movies, err := m.MovieService.FetchMovies()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, movies)
}
