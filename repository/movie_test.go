package repository

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	pg "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/gorm"
	"path/filepath"
	"testing"
	"time"
)

type movieRepositoryTestSuite struct {
	suite.Suite
	movieRepository *DefaultMovieRepository
	db              *gorm.DB
}

func TestMovieRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(movieRepositoryTestSuite))
}

func (s *movieRepositoryTestSuite) SetupSuite() {
	ctx := context.Background()

	pgContainer, err := pg.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		pg.WithInitScripts(filepath.Join("..", "testdata", "schema.sql")),
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
	if err != nil {
		s.T().Fatal(err)
	}
	fmt.Println(connStr)
	gormDB := NewPostgresConnection(connStr)
	s.db = gormDB

	s.movieRepository = &DefaultMovieRepository{
		db: s.db,
	}
}

func (s *movieRepositoryTestSuite) SetupSubTest() {
	s.db.Exec("TRUNCATE TABLE movies")
}

func (s *movieRepositoryTestSuite) TestFindMovies() {
	s.Run("should find movies", func() {
		movieEntity := MovieEntity{
			ID:   99,
			Name: "test movie",
		}
		s.db.Create(&movieEntity)

		movies, err := s.movieRepository.FindMovies()

		assert.NoError(s.T(), err)
		assert.Equal(s.T(), 1, len(movies))
		assert.Equal(s.T(), 99, movies[0].ID)
		assert.Equal(s.T(), "test movie", movies[0].Name)
	})
}
