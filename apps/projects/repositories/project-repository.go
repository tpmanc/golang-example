package repositories

import (
	"errors"
	"github.com/tpmanc/go-projects/models"
	"gorm.io/gorm"
)

type ProjectRepositoryInterface interface {
	GetById(id string) *models.Project
	GetAll() *[]models.Project
	GetAllByUserId(userId int) *[]models.Project
	Create(project *models.Project)
	Update(project *models.Project)
	Delete(id string) bool
}

type projectRepository struct {
	Db *gorm.DB
}

func (rep *projectRepository) GetById(id string) *models.Project {
	var item models.Project
	err := rep.Db.First(&item, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	return &item
}

func (rep *projectRepository) GetAll() *[]models.Project {
	var items []models.Project
	rep.Db.Find(&items)

	return &items
}

func (rep *projectRepository) GetAllByUserId(userId int) *[]models.Project {
	var items []models.Project
	rep.Db.Where("user_id = ?", userId).Find(&items)

	return &items
}

func (rep *projectRepository) Create(project *models.Project) {
	rep.Db.Create(&project)
}

func (rep *projectRepository) Update(project *models.Project) {
	rep.Db.Save(&project)
}

func (rep *projectRepository) Delete(id string) bool {
	res := rep.Db.Delete(&models.Project{}, id)
	if res.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func GetProjectRepository(db *gorm.DB) ProjectRepositoryInterface {
	return &projectRepository{
		Db: db,
	}
}
