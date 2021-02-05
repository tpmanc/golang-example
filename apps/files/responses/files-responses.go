package responses

import "github.com/tpmanc/files/models"

type FilesResponse struct {
	Items *[]models.Files
}

type FileResponse struct {
	Item *models.Files
}

type FilesSaveResponse struct {
	Item *models.Files
}

type FilesDeleteResponse struct {
	Result bool
}
