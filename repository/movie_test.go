package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"path/filepath"
	"testing"
	"time"

	pg "github.com/testcontainers/testcontainers-go/modules/postgres"
)

type movieRepositoryTestSuite struct {
	suite.Suite
	operations *operations
}

func TestMovieRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(movieRepositoryTestSuite))
}

func (s *movieRepositoryTestSuite) SetupSuite() {
	ctx := context.Background()

	pgContainer, err := pg.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		pg.WithInitScripts(filepath.Join("..", "testdata", "init-db.sql")),
		pg.WithDatabase("test-db"),
		pg.WithUsername("postgres"),
		pg.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)

	if err != nil {
		s.T().Fatal(err)
	}

	s.T().Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			s.T().Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	connStr, err := pgContainer.ConnectionString(ctx)
	assert.NoError(s.T(), err)

	s.operations = NewOperations(connStr)
}

func (s *movieRepositoryTestSuite) TestWithTestContainers() {
	movies, err := s.operations.FindMovies()
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 1, len(movies))
	assert.Equal(s.T(), 99, movies[0].ID)
	assert.Equal(s.T(), "test movie", movies[0].Name)
}
