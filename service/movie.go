package service

import (
	"context"
	"log/slog"
	"movie-service/model"
	"movie-service/repository"
)

//go:generate mockery --name MovieService
type MovieService interface {
	FetchMovies(ctx context.Context) ([]model.Movie, error)
}

type DefaultMovieService struct {
	MovieRepository repository.MovieRepository
	Logger          *slog.Logger
}

func NewMovieService(movieRepository repository.MovieRepository, logger *slog.Logger) *DefaultMovieService {
	return &DefaultMovieService{MovieRepository: movieRepository, Logger: logger}
}

func (m *DefaultMovieService) FetchMovies(ctx context.Context) ([]model.Movie, error) {
	fetchedMovies, err := m.MovieRepository.FindMovies()
	if err != nil {
		return nil, err
	}

	var movies []model.Movie
	for _, movie := range fetchedMovies {
		movies = append(movies, model.Movie{
			Id:    movie.ID,
			Title: movie.Name,
		})
	}

	m.Logger.InfoContext(ctx, "Movies returned successfully")

	return movies, nil
}
