package db

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/logger"
	_ "github.com/lib/pq"
)

func NewPostgresDB(pgDsn string) *sqlx.DB {
	db, err := sqlx.Open("postgres", pgDsn)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer db.Close()

	return db
}
