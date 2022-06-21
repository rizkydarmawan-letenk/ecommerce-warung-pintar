package repository

import (
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/models"
	"gorm.io/gorm"
)

type Repository interface {
	Save(user models.User) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user models.User) (models.User, error) {
	// Create new user on db
	err := r.db.Save(&user).Error
	// If err return object data user, with error
	if err != nil {
		return user, err
	}

	// If success return new data user, with no error
	return user, nil
}
