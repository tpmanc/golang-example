package services

import (
	"github.com/tpmanc/go-projects/models"
	"github.com/tpmanc/go-projects/repositories"
	"github.com/tpmanc/go-projects/requests"
)

type ProjectServiceInterface interface {
	GetById(request *requests.ProjectGetByIdRequest) *models.Project
	GetAll(r *requests.ProjectsRequest) *[]models.Project
	Save(request *requests.ProjectSaveRequest) *models.Project
	Delete(request *requests.ProjectDeleteRequest) bool
}

type projectService struct {
	repository repositories.ProjectRepositoryInterface
}

func GetProjectService(repository repositories.ProjectRepositoryInterface) ProjectServiceInterface {
	return &projectService{
		repository: repository,
	}
}

func (s *projectService) GetById(request *requests.ProjectGetByIdRequest) *models.Project {
	return s.repository.GetById(request.Id)
}

func (s *projectService) GetAll(r *requests.ProjectsRequest) *[]models.Project {
	return s.repository.GetAllByUserId(r.UserId)
}

func (s *projectService) Save(request *requests.ProjectSaveRequest) *models.Project {
	var model models.Project

	if len(request.Id) == 0 {
		model.Title = request.Title
		model.UserId = request.UserId
		s.repository.Create(&model)
	} else {
		model = *s.repository.GetById(request.Id)
		model.Title = request.Title
		model.UserId = request.UserId
		s.repository.Update(&model)
	}

	return &model
}

func (s *projectService) Delete(r *requests.ProjectDeleteRequest) bool {
	return s.repository.Delete(r.Id)
}
