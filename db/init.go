package db

import (
	"context"
	"fmt"
	"go-prompt/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s",
		config.GlobalConfig.Database.Host,
		config.GlobalConfig.Database.Port,
		config.GlobalConfig.Database.Db,
		config.GlobalConfig.Database.User,
		config.GlobalConfig.Database.Password,
	)

	dbconfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		// handle error
	}
	dbconfig.AfterConnect = RegisterEnums

	pool, err := pgxpool.NewWithConfig(ctx, dbconfig)
	return pool, nil
}
