package pg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

type queryFunc func(ctx context.Context, sql string, args ...interface{})

func Conn() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	return conn
}

func Execute(sql string, args ...interface{}) {

}
