package repository

import "gorm.io/gorm"

type Repository struct {
	Conn *gorm.DB
}

func New(conn *gorm.DB) *Repository {
	return &Repository{Conn: conn}
}
