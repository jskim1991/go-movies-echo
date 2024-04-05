package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	mocks "movie-service/mocks/movie-service/pkg/repository"
	"movie-service/pkg/repository"
	"testing"
)

type movieServiceTestSuite struct {
	suite.Suite
	mockMovieRepository *mocks.MockMovieRepository
	movieService        *DefaultMovieService
}

func TestMovieServiceTestSuite(t *testing.T) {
	suite.Run(t, new(movieServiceTestSuite))
}

func (s *movieServiceTestSuite) SetupSubTest() {
	s.mockMovieRepository = mocks.NewMockMovieRepository(s.T())
	s.movieService = NewDefaultMovieService(s.mockMovieRepository)
}

func (s *movieServiceTestSuite) TestFetchMovies() {
	s.Run("should call movie repository to fetch movies", func() {
		s.mockMovieRepository.EXPECT().FindMovies().Return([]repository.MovieEntity{}, nil)

		s.movieService.FetchMovies(context.TODO())
	})

	s.Run("should return movies", func() {
		s.mockMovieRepository.EXPECT().FindMovies().Return([]repository.MovieEntity{
			{
				ID:   1,
				Name: "Movie 1",
			},
		}, nil)

		actual, _ := s.movieService.FetchMovies(context.TODO())

		assert.Equal(s.T(), 1, len(actual))
		assert.Equal(s.T(), 1, actual[0].Id)
		assert.Equal(s.T(), "Movie 1", actual[0].Title)
	})

	s.Run("should return error when movie repository returns error", func() {
		s.mockMovieRepository.EXPECT().FindMovies().Return(nil, assert.AnError)

		_, err := s.movieService.FetchMovies(context.TODO())

		assert.Error(s.T(), err)
	})
}
