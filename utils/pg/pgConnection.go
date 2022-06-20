package pg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

func Conn() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:secret@postgres:5432/repaircat")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	return conn
}
