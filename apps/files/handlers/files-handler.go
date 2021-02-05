package handlers

import (
	"github.com/tpmanc/files/db"
	"github.com/tpmanc/files/helpers"
	"github.com/tpmanc/files/repositories"
	"github.com/tpmanc/files/requests"
	"github.com/tpmanc/files/responses"
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
	response := responses.FilesResponse{
		Items: items,
	}
	helpers.ResponseJson(w, response)
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

	response := responses.FileResponse{
		Item: item,
	}
	helpers.ResponseJson(w, response)
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

	response := responses.FilesSaveResponse{
		Item: item,
	}
	helpers.ResponseJson(w, response)
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
		response := responses.FilesDeleteResponse{
			Result: true,
		}
		helpers.ResponseJson(w, response)
	} else {
		helpers.Response500(w, "Delete error")
	}
}
