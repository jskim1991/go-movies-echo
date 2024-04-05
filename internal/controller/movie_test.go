package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	mocks "movie-service/mocks/movie-service/pkg/service"
	"movie-service/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

type movieControllerTestSuite struct {
	suite.Suite
	rec              *httptest.ResponseRecorder
	context          echo.Context
	mockMovieService *mocks.MockMovieService
}

func TestMovieControllerTestSuite(t *testing.T) {
	suite.Run(t, new(movieControllerTestSuite))
}

func (s *movieControllerTestSuite) SetupSubTest() {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/movies", nil)
	s.rec = httptest.NewRecorder()
	s.context = e.NewContext(req, s.rec)

	s.mockMovieService = mocks.NewMockMovieService(s.T())
}

func (s *movieControllerTestSuite) TestGetMovies() {
	s.Run("should return status ok", func() {
		s.mockMovieService.EXPECT().FetchMovies(mock.Anything).Return([]model.Movie{}, nil)
		movieController := NewMovieController(s.mockMovieService)

		err := movieController.GetMovies(s.context)

		assert.Nil(s.T(), err)
		assert.Equal(s.T(), http.StatusOK, s.rec.Code)
	})

	s.Run("should call movie service", func() {
		s.mockMovieService.EXPECT().FetchMovies(mock.Anything).Return([]model.Movie{}, nil)
		movieController := NewMovieController(s.mockMovieService)

		movieController.GetMovies(s.context)
	})

	s.Run("should return movies", func() {
		s.mockMovieService.EXPECT().FetchMovies(mock.Anything).Return([]model.Movie{
			{
				Id:    12,
				Title: "Last Samurai",
			},
		}, nil)
		movieController := NewMovieController(s.mockMovieService)

		movieController.GetMovies(s.context)

		var actual []model.Movie
		json.Unmarshal(s.rec.Body.Bytes(), &actual)

		assert.Equal(s.T(), 1, len(actual))
		assert.Equal(s.T(), 12, actual[0].Id)
		assert.Equal(s.T(), "Last Samurai", actual[0].Title)
	})

	s.Run("should return error when movie service returns error", func() {
		s.mockMovieService.EXPECT().FetchMovies(mock.Anything).Return([]model.Movie{}, assert.AnError)
		movieController := NewMovieController(s.mockMovieService)

		err := movieController.GetMovies(s.context)

		assert.Error(s.T(), err)
	})
}
