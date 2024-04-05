package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"movie-service/mocks"
	"movie-service/repository"
	"testing"
)

type movieServiceTestSuite struct {
	suite.Suite
	mockMovieRepository mocks.MockMovieRepository
	movieService        *DefaultMovieService
}

func TestMovieServiceTestSuite(t *testing.T) {
	suite.Run(t, new(movieServiceTestSuite))
}

func (s *movieServiceTestSuite) SetupTest() {
	s.mockMovieRepository = mocks.MockMovieRepository{}
	s.movieService = NewDefaultMovieService(&s.mockMovieRepository)
}

func (s *movieServiceTestSuite) TestFetchMoviesCallsMovieRepository() {
	s.mockMovieRepository.On("FindMovies").Return([]repository.MovieEntity{}, nil)

	s.movieService.FetchMovies(context.TODO())

	s.mockMovieRepository.AssertNumberOfCalls(s.T(), "FindMovies", 1)
}

func (s *movieServiceTestSuite) TestFetchMoviesReturnsMovies() {
	s.mockMovieRepository.On("FindMovies").Return([]repository.MovieEntity{
		{
			ID:   1,
			Name: "Movie 1",
		},
	}, nil)

	actual, _ := s.movieService.FetchMovies(context.TODO())

	assert.Equal(s.T(), 1, len(actual))
	assert.Equal(s.T(), 1, actual[0].Id)
	assert.Equal(s.T(), "Movie 1", actual[0].Title)
}

func (s *movieServiceTestSuite) TestFetchMoviesReturnsErrors() {
	s.mockMovieRepository.On("FindMovies").Return(nil, assert.AnError)

	_, err := s.movieService.FetchMovies(context.TODO())

	assert.Error(s.T(), err)
}
