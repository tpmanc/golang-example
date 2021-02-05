package services

import (
	"github.com/tpmanc/servers/models"
	"github.com/tpmanc/servers/repositories"
	"github.com/tpmanc/servers/requests"
)

type ServerServiceInterface interface {
	GetAll(r *requests.ServersRequest) *[]models.Server
	GetOne(r *requests.ServerRequest) *models.Server
	Save(r *requests.ServerSaveRequest) *models.Server
	Delete(r *requests.ServerDeleteRequest) bool
}

type serverService struct {
	rep repositories.ServerRepositoryInterface
}

func (s *serverService) GetAll(r *requests.ServersRequest) *[]models.Server {
	return s.rep.GetAllByProjectId(r.ProjectId)
}

func (s *serverService) GetOne(r *requests.ServerRequest) *models.Server {
	return s.rep.GetById(r.Id)
}

func (s *serverService) Save(r *requests.ServerSaveRequest) *models.Server {
	var model models.Server

	if len(r.Id) == 0 {
		model.ProjectId = r.ProjectId
		model.Host = r.Host
		model.User = r.User
		model.Password = r.Password
		model.Port = r.Port
		s.rep.Create(&model)
	} else {
		model = *s.rep.GetById(r.Id)
		model.ProjectId = r.ProjectId
		model.Host = r.Host
		model.User = r.User
		model.Password = r.Password
		model.Port = r.Port
		s.rep.Update(&model)
	}

	return &model
}

func (s *serverService) Delete(r *requests.ServerDeleteRequest) bool {
	return s.rep.Delete(r.Id)
}

func GetServerService(rep repositories.ServerRepositoryInterface) ServerServiceInterface {
	return &serverService{
		rep: rep,
	}
}
