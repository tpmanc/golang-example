package handlers

import (
	"github.com/tpmanc/go-projects/db"
	"github.com/tpmanc/go-projects/helpers"
	"github.com/tpmanc/go-projects/repositories"
	"github.com/tpmanc/go-projects/requests"
	"github.com/tpmanc/go-projects/responses"
	"github.com/tpmanc/go-projects/services"
	"net/http"
)

func getProjectService() services.ProjectServiceInterface {
	conn := db.Get()
	repository := repositories.GetProjectRepository(conn)
	return services.GetProjectService(repository)
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	req, err := requests.ParseProjectsRequest(r)
	if err != nil {
		helpers.Response400(w, err.Error())
		return
	}

	service := getProjectService()
	items := service.GetAll(req)

	response := &responses.ProjectsResponse{
		Items: items,
	}
	helpers.ResponseJson(w, response)
}

func ProjectHandler(w http.ResponseWriter, r *http.Request) {
	request, err := requests.ParseProjectGetByIdRequest(r)
	if err != nil {
		helpers.Response404(w, err.Error())
		return
	}

	service := getProjectService()

	item := service.GetById(request)
	if item == nil {
		helpers.Response404(w, "Project not found")
		return
	}

	response := &responses.ProjectResponse{
		Item: item,
	}
	helpers.ResponseJson(w, response)
}

func ProjectSaveHandler(w http.ResponseWriter, r *http.Request) {
	request, err := requests.ParseProjectSaveRequest(r)
	if err != nil {
		helpers.Response404(w, err.Error())
	}

	service := getProjectService()
	item, err := service.Save(request)
	if err != nil {
		helpers.Response400(w, err.Error())
	}

	response := &responses.ProjectSaveResponse{
		Item: item,
	}
	helpers.ResponseJson(w, response)
}

func ProjectDeleteHandler(w http.ResponseWriter, r *http.Request) {
	request, err := requests.ParseProjectDeleteRequest(r)
	if err != nil {
		helpers.Response404(w, err.Error())
		return
	}

	service := getProjectService()
	result := service.Delete(request)
	if !result {
		helpers.Response500(w, "Server error")
		return
	}

	response := &responses.ProjectDeleteResponse{
		Result: true,
	}
	helpers.ResponseJson(w, response)
}
