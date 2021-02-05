package handlers

import (
	"github.com/tpmanc/files/db"
	"github.com/tpmanc/files/helpers"
	"github.com/tpmanc/files/repositories"
	"github.com/tpmanc/files/requests"
	"github.com/tpmanc/files/services"
	"net/http"
)

func getService() services.FilesServiceInterface {
	repository := repositories.GetFilesRepository(db.Get())
	return services.GetFilesService(repository)
}

func ServersHandler(w http.ResponseWriter, r *http.Request)  {
	req := requests.ParseFilesRequest(r)

	service := getService()

	items := service.GetAll(req)
	helpers.ResponseJson(w, items)
}

func ServerHandler(w http.ResponseWriter, r *http.Request)  {
	req := requests.ParseFileRequest(r)

	service := getService()

	item := service.GetOne(req)
	helpers.ResponseJson(w, item)
}

func ServerSaveHandler(w http.ResponseWriter, r *http.Request)  {
	req := requests.ParseFilesSaveRequest(r)

	service := getService()

	item := service.Save(req)
	helpers.ResponseJson(w, item)
}

func ServerDeleteHandler(w http.ResponseWriter, r *http.Request)  {
	req := requests.ParseFilesDeleteRequest(r)

	service := getService()

	res := service.Delete(req)
	if res {
		helpers.ResponseJson(w, map[string]string{
			"result": "ok",
		})
	} else {
		helpers.Response500(w, "Delete error")
	}
}
