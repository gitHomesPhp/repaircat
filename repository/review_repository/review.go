package review_repository

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/entity"
	"github.com/jackc/pgx/v4"
	"os"
)

func Flush(review *entity.Review) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

}
