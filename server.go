package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"movie-service/controller"
	"movie-service/repository"
	"movie-service/service"
	"net/http"
	"os"
)

func main() {
	e := configureEcho()
	e.Start(":8080")
}

func configureEcho() *echo.Echo {
	handler := &controller.MyHandler{Handler: slog.NewJSONHandler(os.Stdout, nil)}
	logger := slog.New(handler)

	e := echo.New()
	e.Use(middleware.RequestID())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(addCustomContext())
	e.Use(middleware.Logger())

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	dsn := "postgres://postgres:@localhost:5432/movies"
	defaultMovieService := service.NewMovieService(repository.NewOperations(dsn), logger)
	movieController := controller.NewMovieController(defaultMovieService, logger)
	e.GET("/movies", movieController.GetMovies)

	return e
}

func addCustomContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {

			requestId := c.Response().Header().Get(echo.HeaderXRequestID)
			ctx := context.WithValue(context.Background(), "requestId", requestId)

			cc := &controller.CustomContext{
				Context: c,
				Ctx:     ctx,
			}

			return next(cc)
		}
	}
}
