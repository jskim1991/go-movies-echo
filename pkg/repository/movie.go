package repository

import "gorm.io/gorm"

type MovieRepository interface {
	FindMovies() ([]MovieEntity, error)
}

type MovieEntity struct {
	ID   int
	Name string
}

func (MovieEntity) TableName() string {
	return "movies"
}

type DefaultMovieRepository struct {
	db *gorm.DB
}

func NewDefaultMovieRepository(db *gorm.DB) MovieRepository {
	return &DefaultMovieRepository{
		db: db,
	}
}

func (m *DefaultMovieRepository) FindMovies() ([]MovieEntity, error) {
	var movies []MovieEntity
	rows, err := m.db.Find(&movies).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		m.db.ScanRows(rows, &movies)
	}

	return movies, nil
}
