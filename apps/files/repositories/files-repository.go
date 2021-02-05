package repositories

import (
	"errors"
	"github.com/tpmanc/files/models"
	"gorm.io/gorm"
)

type FilesRepositoryInterface interface {
	GetById(id string) *models.Files
	GetAll() *[]models.Files
	GetAllByServerId(serverId int) *[]models.Files
	Create(project *models.Files)
	Update(project *models.Files)
	Delete(id string) bool
}

type filesRepository struct {
	Db *gorm.DB
}

func (rep *filesRepository) GetById(id string) *models.Files {
	var item models.Files
	err := rep.Db.First(&item, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	return &item
}

func (rep *filesRepository) GetAll() *[]models.Files {
	var items []models.Files
	rep.Db.Find(&items)

	return &items
}

func (rep *filesRepository) GetAllByServerId(serverId int) *[]models.Files {
	var items []models.Files
	rep.Db.Where("server_id = ?", serverId).Find(&items)

	return &items
}

func (rep *filesRepository) Create(project *models.Files) {
	rep.Db.Create(&project)
}

func (rep *filesRepository) Update(project *models.Files) {
	rep.Db.Save(&project)
}

func (rep *filesRepository) Delete(id string) bool {
	res := rep.Db.Delete(&models.Files{}, id)
	if res.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func GetFilesRepository(db *gorm.DB) FilesRepositoryInterface {
	return &filesRepository{
		Db: db,
	}
}
