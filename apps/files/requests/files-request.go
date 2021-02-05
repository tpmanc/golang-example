package requests

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type FilesRequest struct {
	ServerId int
}
func ParseFilesRequest(r *http.Request) *FilesRequest {
	serverId, _ := strconv.Atoi(mux.Vars(r)["serverId"])

	req := &FilesRequest{
		ServerId: serverId,
	}

	return req
}

type FileRequest struct {
	Id string
}
func ParseFileRequest(r *http.Request) *FileRequest {
	id := mux.Vars(r)["id"]

	req := &FileRequest{
		Id: id,
	}

	return req
}

type FilesSaveRequest struct {
	Id string
	ServerId int
	Path string
}
func ParseFilesSaveRequest(r *http.Request) *FilesSaveRequest {
	serverId, _ := strconv.Atoi(r.FormValue("serverId"))

	req := &FilesSaveRequest{
		Id: r.FormValue("id"),
		ServerId: serverId,
		Path: r.FormValue("path"),
	}

	return req
}

type FilesDeleteRequest struct {
	Id string
}
func ParseFilesDeleteRequest(r *http.Request) *FilesDeleteRequest {
	req := &FilesDeleteRequest{
		Id: r.FormValue("id"),
	}

	return req
}
