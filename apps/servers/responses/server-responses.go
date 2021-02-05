package responses

import "github.com/tpmanc/servers/models"

type ServersResponse struct {
	Items *[]models.Server
}

type ServerResponse struct {
	Item *models.Server
}

type ServerSaveResponse struct {
	Item *models.Server
}

type ServerDeleteResponse struct {
	Result bool
}
