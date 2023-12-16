package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type integrationTestSuite struct {
	suite.Suite
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(integrationTestSuite))
}

func (s *integrationTestSuite) SetupSuite() {
	go func() {
		err := SetUp().Start("localhost:8080")
		if err != nil {
			s.T().Fail()
		}
	}()
}

func (s *integrationTestSuite) TestHealthEndpoint() {
	resp, err := http.Get("http://localhost:8080/health")
	if err != nil {
		s.T().Fail()
	}

	assert.Equal(s.T(), http.StatusOK, resp.StatusCode)
}

func (s *integrationTestSuite) TestMoviesEndpoint() {
	resp, err := http.Get("http://localhost:8080/movies")
	if err != nil {
		s.T().Fail()
	}

	assert.Equal(s.T(), http.StatusOK, resp.StatusCode)
}
