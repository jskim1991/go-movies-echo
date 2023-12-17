package service

import "movies-service/model"

type MovieService interface {
	FetchMovies() ([]model.Movie, error)
}

type DefaultMovieService struct {
}

func (s DefaultMovieService) FetchMovies() ([]model.Movie, error) {
	return []model.Movie{
		{
			ID:    12121,
			Title: "Batman",
		},
	}, nil
}
