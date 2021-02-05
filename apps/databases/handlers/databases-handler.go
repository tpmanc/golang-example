package handlers

import (
	"github.com/tpmanc/databases/db"
	"github.com/tpmanc/databases/helpers"
	"github.com/tpmanc/databases/repositories"
	"github.com/tpmanc/databases/requests"
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
	helpers.ResponseJson(w, items)
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

	helpers.ResponseJson(w, item)
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

	helpers.ResponseJson(w, item)
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
		helpers.ResponseJson(w, map[string]string{
			"result": "ok",
		})
	} else {
		helpers.Response500(w, "Delete error")
	}
}
