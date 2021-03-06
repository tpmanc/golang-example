package requests

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ProjectsRequest struct {
	UserId int
}
func ParseProjectsRequest(r *http.Request) (*ProjectsRequest, error) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["userId"])

	if userId <= 0 {
		return nil, errors.New("userId is required")
	}

	return &ProjectsRequest{
		UserId: userId,
	}, nil
}

type ProjectGetByIdRequest struct {
	Id string
}
func ParseProjectGetByIdRequest(r *http.Request) (*ProjectGetByIdRequest, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id != "" {
		return nil, errors.New("invalid request")
	}
	return &ProjectGetByIdRequest{Id: id}, nil
}


type ProjectSaveRequest struct {
	Id string
	UserId int
	Title string
}
func ParseProjectSaveRequest(r *http.Request) (*ProjectSaveRequest, error) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	userId, _ := strconv.Atoi(r.FormValue("userId"))

	if len(title) == 0 {
		return nil, errors.New("title is required")
	}
	if userId <= 0 {
		return nil, errors.New("userId is required")
	}

	return &ProjectSaveRequest{
		Id: id,
		Title: title,
		UserId: userId,
	}, nil
}

type ProjectDeleteRequest struct {
	Id string
}
func ParseProjectDeleteRequest(r *http.Request) (*ProjectDeleteRequest, error) {
	id := r.FormValue("id")
	if len(id) == 0 {
		return nil, errors.New("invalid request")
	}

	return &ProjectDeleteRequest{
		Id: id,
	}, nil
}
