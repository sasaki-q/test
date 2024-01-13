package repositories

import (
	"github.com/sasaki-q/test/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(name string) (*models.User, error)
	List() ([]*models.User, error)
	GetById(id uint) (*models.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (r UserRepositoryImpl) Create(name string) (*models.User, error) {
	user := &models.User{Name: name}
	err := r.DB.Save(user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r UserRepositoryImpl) List() ([]*models.User, error) {
	var users []*models.User
	err := r.DB.Find(&users).Error

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r UserRepositoryImpl) GetById(id uint) (*models.User, error) {
	var user *models.User
	err := r.DB.Find(&user).Where("id = ?", id).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}
