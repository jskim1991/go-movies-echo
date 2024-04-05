package service

import (
	"context"
	"movie-service/model"
	"movie-service/repository"
)

//go:generate mockery --name MovieService
type MovieService interface {
	FetchMovies(ctx context.Context) ([]model.Movie, error)
}

type DefaultMovieService struct {
	MovieRepository repository.MovieRepository
}

func NewDefaultMovieService(movieRepository repository.MovieRepository) *DefaultMovieService {
	return &DefaultMovieService{MovieRepository: movieRepository}
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

	return movies, nil
}
