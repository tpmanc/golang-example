package handlers

import (
	"github.com/tpmanc/servers/db"
	"github.com/tpmanc/servers/helpers"
	"github.com/tpmanc/servers/repositories"
	"github.com/tpmanc/servers/requests"
	"github.com/tpmanc/servers/services"
	"net/http"
)

func getService() services.ServerServiceInterface {
	repository := repositories.GetServerRepository(db.Get())
	return services.GetServerService(repository)
}

func ServersHandler(w http.ResponseWriter, r *http.Request)  {
	req := requests.ParseServersRequest(r)

	service := getService()

	items := service.GetAll(req)
	helpers.ResponseJson(w, items)
}

func ServerHandler(w http.ResponseWriter, r *http.Request)  {
	req := requests.ParseServerRequest(r)

	service := getService()

	item := service.GetOne(req)
	helpers.ResponseJson(w, item)
}

func ServerSaveHandler(w http.ResponseWriter, r *http.Request)  {
	req := requests.ParseServerSaveRequest(r)

	service := getService()

	item := service.Save(req)
	helpers.ResponseJson(w, item)
}

func ServerDeleteHandler(w http.ResponseWriter, r *http.Request)  {
	req := requests.ParseServerDeleteRequest(r)

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
