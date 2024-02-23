package pg

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGorm(pg *sql.DB) (*gorm.DB, error) {
	conn := postgres.New(postgres.Config{
		Conn: pg,
	})

	return gorm.Open(conn, &gorm.Config{})
}
