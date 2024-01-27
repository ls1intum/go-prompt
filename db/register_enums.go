package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func RegisterEnums(ctx context.Context, conn *pgx.Conn) error {
	t, err := conn.LoadType(ctx, "course")
	if err != nil {
		return err
	}
	conn.TypeMap().RegisterType(t)

	t, err = conn.LoadType(ctx, "_course")
	if err != nil {
		return err
	}
	conn.TypeMap().RegisterType(t)
	return nil
}
