package repository

import (
	"github.com/valentinusdelvin/savebite-be/internal/domain/entity"
	"gorm.io/gorm"
)

type UserRepositoryItf interface {
	CreateUser(user entity.User) error
	GetUserByEmail(email string) (entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryItf {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user entity.User) error {
	err := r.db.Table("users").Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	err := r.db.Table("users").Where("email = ?", email).First(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
