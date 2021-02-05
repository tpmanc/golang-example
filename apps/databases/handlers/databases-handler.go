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
	req := requests.ParseDatabasesRequest(r)
	service := getService()

	items := service.GetAll(req)
	helpers.ResponseJson(w, items)
}

func DatabaseHandler(w http.ResponseWriter, r *http.Request)  {
	req := requests.ParseDatabaseRequest(r)

	service := getService()

	item := service.GetOne(req)
	helpers.ResponseJson(w, item)
}

func DatabasesSaveHandler(w http.ResponseWriter, r *http.Request)  {
	req := requests.ParseDatabasesSaveRequest(r)

	service := getService()

	item := service.Save(req)
	helpers.ResponseJson(w, item)
}

func DatabasesDeleteHandler(w http.ResponseWriter, r *http.Request)  {
	req := requests.ParseDatabasesDeleteRequest(r)

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
