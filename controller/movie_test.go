package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"movie-service/mocks"
	"movie-service/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

type movieControllerTestSuite struct {
	suite.Suite
	rec     *httptest.ResponseRecorder
	context echo.Context
}

func TestMovieControllerTestSuite(t *testing.T) {
	suite.Run(t, new(movieControllerTestSuite))
}

func (s *movieControllerTestSuite) SetupTest() {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/movies", nil)
	s.rec = httptest.NewRecorder()
	s.context = e.NewContext(req, s.rec)
}

func (s *movieControllerTestSuite) TearDownTest() {
	// after each
}

func (s *movieControllerTestSuite) TestGetMoviesReturnsStatusOk() {
	mockMovieService := mocks.MockMovieService{}
	mockMovieService.On("FetchMovies").Return([]model.Movie{}, nil)
	movieController := MovieController{
		MovieService: &mockMovieService,
	}

	err := movieController.GetMovies(s.context)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), http.StatusOK, s.rec.Code)
}

func (s *movieControllerTestSuite) TestGetMoviesReturnsMovies() {
	mockMovieService := mocks.MockMovieService{}
	mockMovieService.On("FetchMovies").Return([]model.Movie{
		{
			Id:    12,
			Title: "Last Samurai",
		},
	}, nil)
	movieController := MovieController{
		MovieService: &mockMovieService,
	}

	movieController.GetMovies(s.context)

	var actual []model.Movie
	json.Unmarshal(s.rec.Body.Bytes(), &actual)

	assert.Equal(s.T(), 1, len(actual))
	assert.Equal(s.T(), 12, actual[0].Id)
	assert.Equal(s.T(), "Last Samurai", actual[0].Title)
}

func (s *movieControllerTestSuite) TestGetMoviesCallsMovieService() {
	mockMovieService := mocks.MockMovieService{}
	mockMovieService.On("FetchMovies").Return([]model.Movie{}, nil)
	movieController := MovieController{
		MovieService: &mockMovieService,
	}

	movieController.GetMovies(s.context)

	mockMovieService.AssertNumberOfCalls(s.T(), "FetchMovies", 1)
}

func (s *movieControllerTestSuite) TestGetMoviesReturnsErrorWhenMovieServiceReturnsError() {
	mockMovieService := mocks.MockMovieService{}
	mockMovieService.On("FetchMovies").Return([]model.Movie{}, assert.AnError)
	movieController := MovieController{
		MovieService: &mockMovieService,
	}

	err := movieController.GetMovies(s.context)

	assert.Error(s.T(), err)
}