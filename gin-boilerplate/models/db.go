package models

import (
	"fmt"
	"time"

	"gin-boilerplate/config"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

var (
	// DB
	DB *sqlx.DB
)

func init() {
	var err error
	if config.DB.Password == `` {
		config.DB.Password = `null`
	}
	dbStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DB.Host, config.DB.Port, config.DB.User, config.DB.Password, config.DB.Table)
	DB, err = sqlx.Connect(`postgres`, dbStr)
	if err != nil {
		logrus.Fatalln(err)
	}
	DB.SetMaxIdleConns(config.DB.MaxIdleConn)
	DB.SetConnMaxLifetime(2 * time.Minute)
	DB.SetMaxOpenConns(config.DB.MaxConn)
}
