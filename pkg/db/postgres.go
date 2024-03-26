package db

import (
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/kampus-merdeka-standardization/boilerplate-go/pkg/logger"
	_ "github.com/lib/pq"
)

func NewPostgresDB(pgHost, pgUser, pgPassword, pgPort, pgDb string) *sqlx.DB {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", pgHost, pgUser, pgPassword, pgDb, pgPort))
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer db.Close()

	return db
}
