package controller

import (
	"github.com/labstack/echo/v4"
	"movie-service/internal/service"
	"net/http"
)

type MovieController struct {
	MovieService service.MovieService
}

func NewMovieController(movieService service.MovieService) *MovieController {
	return &MovieController{MovieService: movieService}
}

func (m *MovieController) GetMovies(c echo.Context) error {
	ctx := c.Request().Context()
	movies, err := m.MovieService.FetchMovies(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, movies)
}
