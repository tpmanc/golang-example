package services

import (
	"errors"
	"github.com/tpmanc/files/models"
	"github.com/tpmanc/files/repositories"
	"github.com/tpmanc/files/requests"
)

type FilesServiceInterface interface {
	GetAll(r *requests.FilesRequest) *[]models.Files
	GetOne(r *requests.FileRequest) *models.Files
	Save(r *requests.FilesSaveRequest) (*models.Files, error)
	Delete(r *requests.FilesDeleteRequest) bool
}

type filesService struct {
	rep repositories.FilesRepositoryInterface
}

func (s *filesService) GetAll(r *requests.FilesRequest) *[]models.Files {
	return s.rep.GetAllByServerId(r.ServerId)
}

func (s *filesService) GetOne(r *requests.FileRequest) *models.Files {
	return s.rep.GetById(r.Id)
}

func (s *filesService) Save(r *requests.FilesSaveRequest) (*models.Files, error) {
	var model models.Files

	if len(r.Id) == 0 {
		model.ServerId = r.ServerId
		model.Path = r.Path

		isValid, err := model.Validate()
		if isValid {
			s.rep.Create(&model)
		} else {
			return nil, errors.New(err)
		}
	} else {
		model = *s.rep.GetById(r.Id)
		model.ServerId = r.ServerId
		model.Path = r.Path

		isValid, err := model.Validate()
		if isValid {
			s.rep.Update(&model)
		} else {
			return nil, errors.New(err)
		}
	}

	return &model, nil
}

func (s *filesService) Delete(r *requests.FilesDeleteRequest) bool {
	return s.rep.Delete(r.Id)
}

func GetFilesService(rep repositories.FilesRepositoryInterface) FilesServiceInterface {
	return &filesService{
		rep: rep,
	}
}
