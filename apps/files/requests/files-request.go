package requests

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type FilesRequest struct {
	ServerId int
}
func ParseFilesRequest(r *http.Request) (*FilesRequest, error) {
	serverId, _ := strconv.Atoi(mux.Vars(r)["serverId"])

	if serverId <= 0 {
		return nil, errors.New("serverId is required")
	}

	return &FilesRequest{
		ServerId: serverId,
	}, nil
}

type FileRequest struct {
	Id string
}
func ParseFileRequest(r *http.Request) (*FileRequest, error) {
	id := mux.Vars(r)["id"]

	if len(id) == 0 {
		return nil, errors.New("ID is required")
	}

	return &FileRequest{
		Id: id,
	}, nil
}

type FilesSaveRequest struct {
	Id string
	ServerId int
	Path string
}
func ParseFilesSaveRequest(r *http.Request) (*FilesSaveRequest, error) {
	serverId, _ := strconv.Atoi(r.FormValue("serverId"))
	path := r.FormValue("path")

	if serverId <= 0 {
		return nil, errors.New("serverId is required")
	}
	if len(path) < 2 {
		return nil, errors.New("path is required")
	}

	return &FilesSaveRequest{
		Id: r.FormValue("id"),
		ServerId: serverId,
		Path: path,
	}, nil
}

type FilesDeleteRequest struct {
	Id string
}
func ParseFilesDeleteRequest(r *http.Request) (*FilesDeleteRequest, error) {
	id := r.FormValue("id")

	if len(id) == 0 {
		return nil, errors.New("id is required")
	}

	return &FilesDeleteRequest{
		Id: id,
	}, nil
}
