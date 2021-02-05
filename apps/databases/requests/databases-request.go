package requests

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type DatabasesRequest struct {
	ServerId int
}
func ParseDatabasesRequest(r *http.Request) *DatabasesRequest {
	serverId, _ := strconv.Atoi(mux.Vars(r)["serverId"])

	req := &DatabasesRequest{
		ServerId: serverId,
	}

	return req
}

type DatabaseRequest struct {
	Id string
}
func ParseDatabaseRequest(r *http.Request) *DatabaseRequest {
	id := mux.Vars(r)["id"]

	req := &DatabaseRequest{
		Id: id,
	}

	return req
}

type DatabasesSaveRequest struct {
	Id string
	ServerId int
	User string
	Password string
	Database string
}
func ParseDatabasesSaveRequest(r *http.Request) *DatabasesSaveRequest {
	serverId, _ := strconv.Atoi(r.FormValue("serverId"))

	req := &DatabasesSaveRequest{
		Id: r.FormValue("id"),
		ServerId: serverId,
		User: r.FormValue("user"),
		Password: r.FormValue("password"),
		Database: r.FormValue("database"),
	}

	return req
}

type DatabasesDeleteRequest struct {
	Id string
}
func ParseDatabasesDeleteRequest(r *http.Request) *DatabasesDeleteRequest {
	req := &DatabasesDeleteRequest{
		Id: r.FormValue("id"),
	}

	return req
}
