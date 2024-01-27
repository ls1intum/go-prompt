package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var atomic_enums = []string{"application_status", "gender", "language_proficiency", "study_degree", "study_program"}
var compose_enums = []string{"course", "device"}

func RegisterEnums(ctx context.Context, conn *pgx.Conn) error {
	// Register atomic enums
	for _, enum := range atomic_enums {
		t, err := conn.LoadType(ctx, enum)
		if err != nil {
			return err
		}
		conn.TypeMap().RegisterType(t)
	}

	// Register compose enums
	for _, enum := range compose_enums {
		t, err := conn.LoadType(ctx, enum)
		if err != nil {
			return err
		}
		conn.TypeMap().RegisterType(t)

		t, err = conn.LoadType(ctx, "_"+enum)
		if err != nil {
			return err
		}
		conn.TypeMap().RegisterType(t)
	}
	return nil
}
