package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresRepository struct {
	db *gorm.DB
}

func NewOperations(connStr string) *postgresRepository {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &postgresRepository{db: db}
}

func (m *postgresRepository) FindMovies() ([]MovieEntity, error) {
	var movies []MovieEntity
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
