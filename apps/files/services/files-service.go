package services

import (
	"github.com/tpmanc/files/models"
	"github.com/tpmanc/files/repositories"
	"github.com/tpmanc/files/requests"
)

type FilesServiceInterface interface {
	GetAll(r *requests.FilesRequest) *[]models.Files
	GetOne(r *requests.FileRequest) *models.Files
	Save(r *requests.FilesSaveRequest) *models.Files
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

func (s *filesService) Save(r *requests.FilesSaveRequest) *models.Files {
	var model models.Files

	if len(r.Id) == 0 {
		model.ServerId = r.ServerId
		model.Path = r.Path
		s.rep.Create(&model)
	} else {
		model = *s.rep.GetById(r.Id)
		model.ServerId = r.ServerId
		model.Path = r.Path
		s.rep.Update(&model)
	}

	return &model
}

func (s *filesService) Delete(r *requests.FilesDeleteRequest) bool {
	return s.rep.Delete(r.Id)
}

func GetFilesService(rep repositories.FilesRepositoryInterface) FilesServiceInterface {
	return &filesService{
		rep: rep,
	}
}
