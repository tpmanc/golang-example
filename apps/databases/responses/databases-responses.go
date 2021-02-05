package responses

import "github.com/tpmanc/databases/models"

type DatabasesResponse struct {
	Items *[]models.Databases
}

type DatabaseResponse struct {
	Item *models.Databases
}

type DatabasesSaveResponse struct {
	Item *models.Databases
}

type DatabasesDeleteResponse struct {
	Result bool
}
