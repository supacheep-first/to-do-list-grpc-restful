package models

type Task struct {
	ID        int    `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Completed bool   `gorm:"default:false"`
}
