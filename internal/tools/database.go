package tools

import (
	"AvitoIntership/config"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
)

func PostgreConnect(c config.RepositoryConfig) *sql.DB {
	connectString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", c.Host, c.Port, c.User, c.Password, c.Database)
	db, err := sql.Open("postgres", connectString)
	if err != nil {
		logrus.Fatal(err)
	}
	return db
}
