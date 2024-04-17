package db

import (
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/kampus-merdeka-standardization/boilerplate-pinger-app/pkg/logger"
	_ "github.com/lib/pq"
)

func NewPostgresDB(pgDsn PostgresDsn) *sqlx.DB {
	db, err := sqlx.Open("postgres", pgDsn.ToString())
	if err != nil {
		logger.Fatal(err.Error())
	}

	return db
}

type PostgresDsn struct {
	Host     string
	Port     int
	User     string
	Password string
	Db       string
}

func (p PostgresDsn) ToString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", p.Host, p.User, p.Password, p.Db, p.Port)
}
