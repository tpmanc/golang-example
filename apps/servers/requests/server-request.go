package requests

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ServersRequest struct {
	ProjectId int
}
func ParseServersRequest(r *http.Request) (*ServersRequest, error) {
	projectId, _ := strconv.Atoi(mux.Vars(r)["projectId"])

	if projectId <= 0 {
		return nil, errors.New("projectId is required")
	}

	req := &ServersRequest{
		ProjectId: projectId,
	}

	return req, nil
}

type ServerRequest struct {
	Id string
}
func ParseServerRequest(r *http.Request) (*ServerRequest, error) {
	id := mux.Vars(r)["id"]

	if id == "" {
		return nil, errors.New("server id is required")
	}

	req := &ServerRequest{
		Id: id,
	}

	return req, nil
}

type ServerSaveRequest struct {
	Id string
	ProjectId int
	Host string
	User string
	Password string
	Port string
}
func ParseServerSaveRequest(r *http.Request) (*ServerSaveRequest, error) {
	projectId, _ := strconv.Atoi(r.FormValue("projectId"))
	host := r.FormValue("host")
	user := r.FormValue("user")
	password := r.FormValue("password")
	port := r.FormValue("port")

	if projectId <= 0 {
		return nil, errors.New("projectId is required")
	}
	if len(host) < 2 {
		return nil, errors.New("host is required")
	}
	if len(user) < 2 {
		return nil, errors.New("user is required")
	}
	if len(password) < 2 {
		return nil, errors.New("password is required")
	}
	if len(port) < 2 {
		return nil, errors.New("port is required")
	}

	req := &ServerSaveRequest{
		Id: r.FormValue("id"),
		ProjectId: projectId,
		Host: host,
		User: user,
		Password: password,
		Port: port,
	}

	return req, nil
}

type ServerDeleteRequest struct {
	Id string
}
func ParseServerDeleteRequest(r *http.Request) (*ServerDeleteRequest, error) {
	id := r.FormValue("id")
	if id == "" {
		return nil, errors.New("server id is required")
	}

	req := &ServerDeleteRequest{
		Id: id,
	}

	return req, nil
}
