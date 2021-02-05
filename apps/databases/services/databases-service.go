package services

import (
	"errors"
	"github.com/tpmanc/databases/models"
	"github.com/tpmanc/databases/repositories"
	"github.com/tpmanc/databases/requests"
)

type DatabasesServiceInterface interface {
	GetAll(r *requests.DatabasesRequest) *[]models.Databases
	GetOne(r *requests.DatabaseRequest) *models.Databases
	Save(r *requests.DatabasesSaveRequest) (*models.Databases, error)
	Delete(r *requests.DatabasesDeleteRequest) bool
}

type databasesService struct {
	rep repositories.DatabasesRepositoryInterface
}

func (s *databasesService) GetAll(r *requests.DatabasesRequest) *[]models.Databases {
	return s.rep.GetAllByServerId(r.ServerId)
}

func (s *databasesService) GetOne(r *requests.DatabaseRequest) *models.Databases {
	return s.rep.GetById(r.Id)
}

func (s *databasesService) Save(r *requests.DatabasesSaveRequest) (*models.Databases, error) {
	var model models.Databases

	if len(r.Id) == 0 {
		model.ServerId = r.ServerId
		model.User = r.User
		model.Password = r.Password
		model.Database = r.Database

		isValid, err := model.Validate()
		if isValid {
			s.rep.Create(&model)
		} else {
			return nil, errors.New(err)
		}
	} else {
		model = *s.rep.GetById(r.Id)
		model.ServerId = r.ServerId
		model.User = r.User
		model.Password = r.Password
		model.Database = r.Database

		isValid, err := model.Validate()
		if isValid {
			s.rep.Update(&model)
		} else {
			return nil, errors.New(err)
		}
	}

	return &model, nil
}

func (s *databasesService) Delete(r *requests.DatabasesDeleteRequest) bool {
	return s.rep.Delete(r.Id)
}

func GetDatabasesService(rep repositories.DatabasesRepositoryInterface) DatabasesServiceInterface {
	return &databasesService{
		rep: rep,
	}
}
