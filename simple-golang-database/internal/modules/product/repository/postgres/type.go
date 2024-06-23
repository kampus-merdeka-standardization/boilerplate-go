package product_postgres

import "github.com/jmoiron/sqlx"

type productPostgresRepository struct {
	db *sqlx.DB
}
