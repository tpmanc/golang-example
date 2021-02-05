package repositories

import (
	"errors"
	"github.com/tpmanc/databases/models"
	"gorm.io/gorm"
)

type DatabasesRepositoryInterface interface {
	GetById(id string) *models.Databases
	GetAll() *[]models.Databases
	GetAllByServerId(serverId int) *[]models.Databases
	Create(project *models.Databases)
	Update(project *models.Databases)
	Delete(id string) bool
}

type databasesRepository struct {
	Db *gorm.DB
}

func (rep *databasesRepository) GetById(id string) *models.Databases {
	var item models.Databases
	err := rep.Db.First(&item, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	return &item
}

func (rep *databasesRepository) GetAll() *[]models.Databases {
	var items []models.Databases
	rep.Db.Find(&items)

	return &items
}

func (rep *databasesRepository) GetAllByServerId(serverId int) *[]models.Databases {
	var items []models.Databases
	rep.Db.Where("server_id = ?", serverId).Find(&items)

	return &items
}

func (rep *databasesRepository) Create(project *models.Databases) {
	rep.Db.Create(&project)
}

func (rep *databasesRepository) Update(project *models.Databases) {
	rep.Db.Save(&project)
}

func (rep *databasesRepository) Delete(id string) bool {
	res := rep.Db.Delete(&models.Databases{}, id)
	if res.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func GetDatabasesRepository(db *gorm.DB) DatabasesRepositoryInterface {
	return &databasesRepository{
		Db: db,
	}
}
