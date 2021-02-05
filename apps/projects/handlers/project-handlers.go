package handlers

import (
	"github.com/tpmanc/go-projects/db"
	"github.com/tpmanc/go-projects/helpers"
	"github.com/tpmanc/go-projects/repositories"
	"github.com/tpmanc/go-projects/requests"
	"github.com/tpmanc/go-projects/services"
	"net/http"
)

func getProjectService() services.ProjectServiceInterface {
	conn := db.Get()
	repository := repositories.GetProjectRepository(conn)
	return services.GetProjectService(repository)
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	req := requests.ParseProjectsRequest(r)

	service := getProjectService()
	items := service.GetAll(req)

	helpers.ResponseJson(w, items)
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

	helpers.ResponseJson(w, item)
}

func ProjectSaveHandler(w http.ResponseWriter, r *http.Request) {
	request, err := requests.ParseProjectSaveRequest(r)
	if err != nil {
		helpers.Response404(w, err.Error())
	}

	service := getProjectService()
	item := service.Save(request)

	helpers.ResponseJson(w, item)
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

	helpers.ResponseJson(w, map[string]bool{
		"result": true,
	})
}
