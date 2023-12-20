package repository

import "movie-service/data"

type MovieRepository interface {
	FindMovies() ([]data.MovieEntity, error)
}
