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
	req, err := requests.ParseServersRequest(r)
	if err != nil {
		helpers.Response400(w, err.Error())
		return
	}

	service := getService()

	items := service.GetAll(req)
	helpers.ResponseJson(w, items)
}

func ServerHandler(w http.ResponseWriter, r *http.Request)  {
	req, err := requests.ParseServerRequest(r)
	if err != nil {
		helpers.Response400(w, err.Error())
		return
	}

	service := getService()

	item := service.GetOne(req)
	if item == nil {
		helpers.Response404(w, "Server not found")
		return
	}

	helpers.ResponseJson(w, item)
}

func ServerSaveHandler(w http.ResponseWriter, r *http.Request)  {
	req, err := requests.ParseServerSaveRequest(r)
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
	req, err := requests.ParseServerDeleteRequest(r)
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
