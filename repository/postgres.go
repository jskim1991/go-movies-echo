package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"movie-service/data"
)

type operations struct {
	db *gorm.DB
}

func NewOperations(connStr string) *operations {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &operations{db: db}
}

func (m *operations) FindMovies() ([]data.MovieEntity, error) {
	var movies []data.MovieEntity
	rows, err := m.db.Raw("SELECT id, name FROM movies").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		m.db.ScanRows(rows, &movies)
	}

	return movies, nil
}
