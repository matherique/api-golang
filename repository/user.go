package repository

import (
	"database/sql"

	"github.com/matherique/api-go/domain"
)

type UserRepository struct {
	database *sql.DB
}

func (repo *UserRepository) List() ([]domain.User, error) {
	return []domain.User{}, nil
}