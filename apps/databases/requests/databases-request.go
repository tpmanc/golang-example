package requests

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type DatabasesRequest struct {
	ServerId int
}
func ParseDatabasesRequest(r *http.Request) (*DatabasesRequest, error) {
	serverId, _ := strconv.Atoi(mux.Vars(r)["serverId"])
	if serverId <= 0 {
		return nil, errors.New("serverId is required")
	}

	return &DatabasesRequest{
		ServerId: serverId,
	}, nil
}

type DatabaseRequest struct {
	Id string
}
func ParseDatabaseRequest(r *http.Request) (*DatabaseRequest, error) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		return nil, errors.New("ID is required")
	}

	return &DatabaseRequest{
		Id: id,
	}, nil
}

type DatabasesSaveRequest struct {
	Id string
	ServerId int
	User string
	Password string
	Database string
}
func ParseDatabasesSaveRequest(r *http.Request) (*DatabasesSaveRequest, error) {
	serverId, _ := strconv.Atoi(r.FormValue("serverId"))
	user := r.FormValue("user")
	password := r.FormValue("password")
	database := r.FormValue("database")

	if serverId <= 0 {
		return nil, errors.New("serverId is required")
	}
	if len(user) <= 2 {
		return nil, errors.New("user is required")
	}
	if len(password) <= 2 {
		return nil, errors.New("password is required")
	}
	if len(database) <= 2 {
		return nil, errors.New("database is required")
	}

	return &DatabasesSaveRequest{
		Id: r.FormValue("id"),
		ServerId: serverId,
		User: user,
		Password: password,
		Database: database,
	}, nil
}

type DatabasesDeleteRequest struct {
	Id string
}
func ParseDatabasesDeleteRequest(r *http.Request) (*DatabasesDeleteRequest, error) {
	id := r.FormValue("id")
	if len(id) == 0 {
		return nil, errors.New("id is required")
	}

	return &DatabasesDeleteRequest{
		Id: id,
	}, nil
}
