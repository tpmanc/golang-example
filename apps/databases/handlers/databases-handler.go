package handlers

import (
	"github.com/tpmanc/databases/db"
	"github.com/tpmanc/databases/helpers"
	"github.com/tpmanc/databases/repositories"
	"github.com/tpmanc/databases/requests"
	"github.com/tpmanc/databases/responses"
	"github.com/tpmanc/databases/services"
	"net/http"
)

func getService() services.DatabasesServiceInterface {
	repository := repositories.GetDatabasesRepository(db.Get())
	return services.GetDatabasesService(repository)
}

func DatabasesHandler(w http.ResponseWriter, r *http.Request)  {
	req, err := requests.ParseDatabasesRequest(r)
	if err != nil {
		helpers.Response400(w, err.Error())
		return
	}

	service := getService()

	items := service.GetAll(req)
	response := responses.DatabasesResponse{Items: items}
	helpers.ResponseJson(w, response)
}

func DatabaseHandler(w http.ResponseWriter, r *http.Request)  {
	req, err := requests.ParseDatabaseRequest(r)
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

	response := responses.DatabaseResponse{Item: item}
	helpers.ResponseJson(w, response)
}

func DatabasesSaveHandler(w http.ResponseWriter, r *http.Request)  {
	req, err := requests.ParseDatabasesSaveRequest(r)
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

	response := responses.DatabasesSaveResponse{Item: item}
	helpers.ResponseJson(w, response)
}

func DatabasesDeleteHandler(w http.ResponseWriter, r *http.Request)  {
	req, err := requests.ParseDatabasesDeleteRequest(r)
	if err != nil {
		helpers.Response400(w, err.Error())
		return
	}

	service := getService()

	res := service.Delete(req)
	if res {
		response := responses.DatabasesDeleteResponse{Result: true}
		helpers.ResponseJson(w, response)
	} else {
		helpers.Response500(w, "Delete error")
	}
}
