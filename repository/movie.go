package repository

//go:generate mockery --name MovieRepository
type MovieRepository interface {
	FindMovies() ([]MovieEntity, error)
}

type MovieEntity struct {
	ID   int `gorm:"primaryKey"`
	Name string
}
