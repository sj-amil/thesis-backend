package models

type Blog struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Title   string `json:"title" gorm:"not null;column:title;size:255"`
	Contnet string `json:"contnet" gorm:"not null;colmn:content;size:255"`
}