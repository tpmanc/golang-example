package requests

import (
	"errors"
	"net/http"
)

type RunRequest struct {
	SshHost string
	SshUser string
	SshPassword string
	SshPort string
	DbUser string
	DbPassword string
	DbName string
}

func (r RunRequest) Validate() bool {
	return true
}

func ParseRunRequest(r *http.Request) (*RunRequest, error) {
	req := &RunRequest{
		SshHost: r.FormValue("sshHost"),
		SshUser: r.FormValue("sshUser"),
		SshPassword: r.FormValue("sshPassword"),
		SshPort: r.FormValue("sshPort"),
		DbUser: r.FormValue("dbUser"),
		DbPassword: r.FormValue("dbPassword"),
		DbName: r.FormValue("dbName"),
	}

	if req.Validate() {
		return req, nil
	} else {
		return nil, errors.New("invalid request")
	}
}
