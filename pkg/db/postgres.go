package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/logger"
	_ "github.com/lib/pq"
)

func NewPostgresDB(connString string) *sqlx.DB {
	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		logger.Fatal(err.Error())
	}

	return db
}
