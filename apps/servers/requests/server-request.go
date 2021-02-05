package requests

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ServersRequest struct {
	ProjectId int
}
func ParseServersRequest(r *http.Request) *ServersRequest {
	projectId, _ := strconv.Atoi(mux.Vars(r)["projectId"])
	req := &ServersRequest{
		ProjectId: projectId,
	}

	return req
}

type ServerRequest struct {
	Id string
}
func ParseServerRequest(r *http.Request) *ServerRequest {
	id := mux.Vars(r)["id"]

	req := &ServerRequest{
		Id: id,
	}

	return req
}

type ServerSaveRequest struct {
	Id string
	ProjectId int
	Host string
	User string
	Password string
	Port string
}
func ParseServerSaveRequest(r *http.Request) *ServerSaveRequest {
	projectId, _ := strconv.Atoi(r.FormValue("projectId"))
	req := &ServerSaveRequest{
		Id: r.FormValue("id"),
		ProjectId: projectId,
		Host: r.FormValue("host"),
		User: r.FormValue("user"),
		Password: r.FormValue("password"),
		Port: r.FormValue("port"),
	}

	return req
}

type ServerDeleteRequest struct {
	Id string
}
func ParseServerDeleteRequest(r *http.Request) *ServerDeleteRequest {
	req := &ServerDeleteRequest{
		Id: r.FormValue("id"),
	}

	return req
}
