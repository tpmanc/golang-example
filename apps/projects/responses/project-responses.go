package responses

import "github.com/tpmanc/go-projects/models"

type ProjectsResponse struct {
	Items *[]models.Project
}

type ProjectResponse struct {
	Item *models.Project
}

type ProjectSaveResponse struct {
	Item *models.Project
}

type ProjectDeleteResponse struct {
	Result bool
}
