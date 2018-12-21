package mongo

import (
	"database/sql"

	"github.com/mattrax/Mattrax/internal/mattrax"
)

type userRepository struct {
	db *sql.DB
}

func (r userRepository) Create(user *mattrax.User) error {
	return nil
}

func (r userRepository) Remove(user *mattrax.User) error {
	return nil

}

func (r userRepository) Find(id mattrax.UserID) (*mattrax.User, error) {
	return &mattrax.User{}, nil
}

func NewUserRepository(db *sql.DB) mattrax.UserRepository {
	return userRepository{db}
}
