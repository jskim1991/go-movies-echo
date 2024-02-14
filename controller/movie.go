package controller

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"movie-service/service"
	"net/http"
)

type MovieController struct {
	MovieService service.MovieService
	logger       *slog.Logger
}

func NewMovieController(movieService service.MovieService, logger *slog.Logger) *MovieController {
	return &MovieController{MovieService: movieService, logger: logger}
}

func (m *MovieController) GetMovies(c echo.Context) error {
	cc := c.(*CustomContext)
	movies, err := m.MovieService.FetchMovies(cc.Ctx)
	if err != nil {
		return err
	}

	m.logger.InfoContext(cc.Ctx, "Movies fetched successfully")

	return c.JSON(http.StatusOK, movies)
}
