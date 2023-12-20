package data

type MovieEntity struct {
	ID   int `gorm:"primaryKey"`
	Name string
}
