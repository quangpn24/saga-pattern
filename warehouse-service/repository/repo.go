package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
}

func New(db *gorm.DB) *Repository {
	return &Repository{}
}
