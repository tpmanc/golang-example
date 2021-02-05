package databases

import "fmt"

type DatabaseBackuperInterface interface {
	// todo: input params
	GetCommand(dbUser string, dbPassword string, dbName string) string
}

type mysqlBackuper struct {}
func (db *mysqlBackuper) GetCommand(dbUser string, dbPassword string, dbName string) string {
	cmd := fmt.Sprintf("mysqldump -u %s -p%s %s > /tmp/go-dump.sql", dbUser, dbPassword, dbName)
	//cmd := "mysqldump -u root -pPa$$w0rd kesha_sugar > /tmp/go-dump.sql"
	fmt.Println(cmd)
	return cmd
}

type postgresBackuper struct {}
func (db *postgresBackuper) GetCommand(dbUser string, dbPassword string, dbName string) string {
	return "start postgres backup"
}

func GetBackuper(dbType string) DatabaseBackuperInterface {
	switch dbType {
		case "mysql":
			return &mysqlBackuper{}
		case "postgres":
			return &postgresBackuper{}
	}

	return nil
}
