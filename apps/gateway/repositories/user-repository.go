package repositories

import (
	"github.com/tpmanc/gateway/models"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetUserByLogin(login string) *models.User
	Create(user *models.User)
	Update(user *models.User)
}

type userRepository struct {
	Db *gorm.DB
}

func (r *userRepository) GetUserByLogin(login string) *models.User {
	var model models.User
	err := r.Db.Where(&models.User{Username: login}).First(&model).Error
	if err != nil {
		return nil
	}

	return &model
}

func (r *userRepository) Create(user *models.User) {
	r.Db.Create(&user)
}

func (r *userRepository) Update(user *models.User) {
	r.Db.Save(&user)
}

func GetUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{
		Db: db,
	}
}
