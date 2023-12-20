package mocks

import (
	"github.com/stretchr/testify/mock"
	"movie-service/data"
	"movie-service/model"
)

type MockMovieService struct {
	mock.Mock
}

func (m *MockMovieService) FetchMovies() ([]model.Movie, error) {
	args := m.Called()
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}
	return args.Get(0).([]model.Movie), nil
}

type MockMovieRepository struct {
	mock.Mock
}

func (m *MockMovieRepository) FindMovies() ([]data.MovieEntity, error) {
	args := m.Called()
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}
	return args.Get(0).([]data.MovieEntity), nil
}
