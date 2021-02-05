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
	req, err := requests.ParseFilesRequest(r)
	if err != nil {
		helpers.Response400(w, err.Error())
		return
	}

	service := getService()

	items := service.GetAll(req)
	helpers.ResponseJson(w, items)
}

func ServerHandler(w http.ResponseWriter, r *http.Request)  {
	req, err := requests.ParseFileRequest(r)
	if err != nil {
		helpers.Response400(w, err.Error())
		return
	}

	service := getService()

	item := service.GetOne(req)
	if item == nil {
		helpers.Response404(w, "Item not found")
		return
	}

	helpers.ResponseJson(w, item)
}

func ServerSaveHandler(w http.ResponseWriter, r *http.Request)  {
	req, err := requests.ParseFilesSaveRequest(r)
	if err != nil {
		helpers.Response400(w, err.Error())
		return
	}

	service := getService()

	item, err := service.Save(req)
	if err != nil {
		helpers.Response400(w, err.Error())
		return
	}

	helpers.ResponseJson(w, item)
}

func ServerDeleteHandler(w http.ResponseWriter, r *http.Request)  {
	req, err := requests.ParseFilesDeleteRequest(r)
	if err != nil {
		helpers.Response400(w, err.Error())
		return
	}

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
