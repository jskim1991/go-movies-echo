package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"movie-service/model"
	"movie-service/repository"
)

type MockMovieService struct {
	mock.Mock
}

func (m *MockMovieService) FetchMovies(ctx context.Context) ([]model.Movie, error) {
	args := m.Called()
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}
	return args.Get(0).([]model.Movie), nil
}

type MockMovieRepository struct {
	mock.Mock
}

func (m *MockMovieRepository) FindMovies() ([]repository.MovieEntity, error) {
	args := m.Called()
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}
	return args.Get(0).([]repository.MovieEntity), nil
}
