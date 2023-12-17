package mocks

import (
	"github.com/stretchr/testify/mock"
	"movies-service/model"
)

type MockMovieService struct {
	mock.Mock
}

func (m *MockMovieService) FetchMovies() ([]model.Movie, error) {
	args := m.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Movie), nil
}
