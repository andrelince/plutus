package pg

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Postgres struct {
	db *sql.DB
}

type PostgresSettings struct {
	MaxOpenConns int    `env:"POSTGRES_MAX_OPEN"`
	MaxIdleConns int    `env:"POSTGRES_MAX_IDLE"`
	DataSource   string `env:"POSTGRES_URL"`
}

func NewPostgres(config PostgresSettings) (*Postgres, error) {
	db, err := sql.Open("pgx", config.DataSource)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)

	return &Postgres{db}, nil
}

func (p Postgres) Close() error {
	return p.db.Close()
}

func (p Postgres) DB() *sql.DB {
	return p.db
}
