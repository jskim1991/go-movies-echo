package service

import (
	"movie-service/model"
	"movie-service/repository"
)

type MovieService interface {
	FetchMovies() ([]model.Movie, error)
}

type DefaultMovieService struct {
	MovieRepository repository.MovieRepository
}

func NewMovieService(movieRepository repository.MovieRepository) *DefaultMovieService {
	return &DefaultMovieService{MovieRepository: movieRepository}
}

func (m *DefaultMovieService) FetchMovies() ([]model.Movie, error) {
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
