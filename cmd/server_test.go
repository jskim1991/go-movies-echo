package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type serverTestSuite struct {
	suite.Suite
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(serverTestSuite))
}

func (s *serverTestSuite) SetupSuite() {
	e := configureEcho()

	go func() {
		err := e.Start(":8080")
		panic(err)
	}()
}

func (s *serverTestSuite) TestHealthEndpoint() {
	resp, err := http.Get("http://localhost:8080/health")

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), http.StatusOK, resp.StatusCode)
}

func (s *serverTestSuite) TestMoviesEndpoint() {
	resp, _ := http.Get("http://localhost:8080/movies")

	assert.Equal(s.T(), http.StatusOK, resp.StatusCode)
}
