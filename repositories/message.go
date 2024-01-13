package repositories

import (
	"github.com/sasaki-q/test/models"
	"gorm.io/gorm"
)

type MessageRepository interface {
	Create(title string, userId uint) (*models.Message, error)
	List() ([]*models.Message, error)
	GetById(id uint) (*models.Message, error)
}

type MessageRepositoryImpl struct {
	DB *gorm.DB
}

func (r MessageRepositoryImpl) Create(title string, userId uint) (*models.Message, error) {
	msg := &models.Message{Title: title, UserID: userId}
	err := r.DB.Save(msg).Error

	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (r MessageRepositoryImpl) List() ([]*models.Message, error) {
	var msgs []*models.Message
	err := r.DB.Find(&msgs).Error

	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func (r MessageRepositoryImpl) GetById(id uint) (*models.Message, error) {
	var msg *models.Message
	err := r.DB.Find(&msg).Where("id = ?", id).Error

	if err != nil {
		return nil, err
	}
	return msg, nil
}
