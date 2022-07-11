package sc_repository

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/entity"
	"github.com/jackc/pgx/v4"
	"os"
)

func Flush(sc *entity.Sc) error {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	_, err = conn.Exec(context.Background(), InsertSc,
		sc.GetName(),
		sc.GetDescription(),
		sc.GetPhone(),
		sc.GetEmail(),
		sc.GetSite(),
		sc.GetLocation().GetId(),
	)

	if err != nil {
		return err
	}

	defer conn.Close(context.Background())

	return nil
}
