package repository

import (
	"fmt"
	"go-with-kafka/internal/model"

	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

type UserRepository interface {
	Create(entity model.UserModel) (*model.UserModel, error)
}

func NewUserRepositoryDB(db *gorm.DB) UserRepository {
	return userRepositoryDB{db}
}

func (r userRepositoryDB) Create(entity model.UserModel) (*model.UserModel, error) {

	entityOriginal := model.UserModel{}
	err := r.db.Where("username = ?", entity.Username).First(&entityOriginal).Error
	if err == nil {
		return nil, fmt.Errorf("user with username %s already exists", entity.Username)
	}
	if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if tx := r.db.Create(&entity); tx.Error != nil {
		return nil, tx.Error
	}
	return &entity, nil
}
