package handlers

import (
	"github.com/tpmanc/backup-db/helpers"
	"github.com/tpmanc/backup-db/requests"
	"github.com/tpmanc/backup-db/services"
	"net/http"
)

func RunHandler(w http.ResponseWriter, r *http.Request)  {
	req, err := requests.ParseRunRequest(r)
	if err != nil {
		helpers.Response500(w, err.Error())
		return
	}

	service := services.GetRunService()
	service.RunBackup(req)

	helpers.ResponseJson(w, map[string]bool{
		"result": true,
	})
}