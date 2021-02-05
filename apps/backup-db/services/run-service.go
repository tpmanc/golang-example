package services

import (
	"fmt"
	"github.com/tpmanc/backup-db/components"
	"github.com/tpmanc/backup-db/databases"
	"github.com/tpmanc/backup-db/requests"
)

type RunServiceInterface interface {
	RunBackup(r *requests.RunRequest)
}

type runService struct {}

func (s *runService) RunBackup(r *requests.RunRequest) {
	fmt.Println("run db backup")

	sshConnect := components.GetSshConnect(r.SshHost, r.SshUser, r.SshPassword, r.SshPort)
	backuper := databases.GetBackuper("mysql")
	sshConnect.RunCommand(backuper.GetCommand(r.DbUser, r.DbPassword, r.DbName))
	fmt.Println("finish backup")
	sshConnect.RunCommand("ls /")
	sshConnect.DownloadFile("/tmp/dump.sql", "/Users/tpmanc/Sites/dump.sql")
	sshConnect.Close()
}

func GetRunService() RunServiceInterface {
	return &runService{}
}
