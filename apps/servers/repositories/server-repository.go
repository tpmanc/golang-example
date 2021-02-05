package repositories

import (
	"errors"
	"github.com/tpmanc/servers/models"
	"gorm.io/gorm"
)

type ServerRepositoryInterface interface {
	GetById(id string) *models.Server
	GetAll() *[]models.Server
	GetAllByProjectId(projectId int) *[]models.Server
	Create(project *models.Server)
	Update(project *models.Server)
	Delete(id string) bool
}

type projectRepository struct {
	Db *gorm.DB
}

func (rep *projectRepository) GetById(id string) *models.Server {
	var item models.Server
	err := rep.Db.First(&item, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	return &item
}

func (rep *projectRepository) GetAll() *[]models.Server {
	var items []models.Server
	rep.Db.Find(&items)

	return &items
}

func (rep *projectRepository) GetAllByProjectId(projectId int) *[]models.Server {
	var items []models.Server
	rep.Db.Where("project_id = ?", projectId).Find(&items)

	return &items
}

func (rep *projectRepository) Create(project *models.Server) {
	rep.Db.Create(&project)
}

func (rep *projectRepository) Update(project *models.Server) {
	rep.Db.Save(&project)
}

func (rep *projectRepository) Delete(id string) bool {
	res := rep.Db.Delete(&models.Server{}, id)
	if res.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func GetServerRepository(db *gorm.DB) ServerRepositoryInterface {
	return &projectRepository{
		Db: db,
	}
}
