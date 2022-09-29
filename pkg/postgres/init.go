package postgres

import (
	"BalanceUser/config"
	"fmt"
	"context"

	"github.com/jmoiron/sqlx"
)

const (
	queryCreateTable = `CREATE TABLE IF NOT EXISTS balance_user (
		id SERIAL PRIMARY KEY,
		user INTEGER,
		balance decimal
	  );`
	queryInsert = `INSERT INTO social_connections ( sender, recipient, number_of_communications)
					VALUES ($1, $2, 0);`
)

func InitPsqlDB(ctx context.Context, cfg *config.Config) (*sqlx.DB, error) {
	connectionURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode,
	)

	database, err := sqlx.Open("postgres", connectionURL)
	if err != nil {
		return nil, err
	}

	if err = database.Ping(); err != nil {
		return nil, err
	}

	database.MustExec(queryCreateTable)

	return database, nil
}