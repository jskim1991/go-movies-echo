package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"movies-service/mocks"
	"movies-service/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

type movieControllerSuite struct {
	suite.Suite
	rec              *httptest.ResponseRecorder
	context          echo.Context
	mockMovieService mocks.MockMovieService
	controller       MovieController
}

func TestMovieControllerSuite(t *testing.T) {
	suite.Run(t, new(movieControllerSuite))
}

func (s *movieControllerSuite) SetupTest() {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/movies", nil)
	s.rec = httptest.NewRecorder()
	s.context = e.NewContext(req, s.rec)

	s.mockMovieService = mocks.MockMovieService{}
	s.controller = MovieController{
		MovieService: &s.mockMovieService,
	}
}

func (s *movieControllerSuite) TestGetMoviesReturnsStatusOk() {
	s.mockMovieService.On("FetchMovies").Return([]model.Movie{}, nil)

	s.controller.GetMovies(s.context)

	assert.Equal(s.T(), http.StatusOK, s.rec.Code)
}

func (s *movieControllerSuite) TestGetMoviesCallsMovieService() {
	s.mockMovieService.On("FetchMovies").Return([]model.Movie{}, nil)

	s.controller.GetMovies(s.context)

	s.mockMovieService.AssertExpectations(s.T())
}

func (s *movieControllerSuite) TestGetMoviesReturnsMovies() {
	s.mockMovieService.On("FetchMovies").Return([]model.Movie{
		{
			ID:    1,
			Title: "Spiderman",
		},
	}, nil)

	s.controller.GetMovies(s.context)

	var actual []model.Movie
	json.Unmarshal(s.rec.Body.Bytes(), &actual)

	assert.Equal(s.T(), 1, len(actual))
	assert.Equal(s.T(), 1, actual[0].ID)
	assert.Equal(s.T(), "Spiderman", actual[0].Title)
}

func (s *movieControllerSuite) TestGetMoviesReturnsErrorWhenMovieServiceReturnsError() {
	s.mockMovieService.On("FetchMovies").Return(new([]model.Movie), assert.AnError)

	err := s.controller.GetMovies(s.context)

	assert.Error(s.T(), err)
}
