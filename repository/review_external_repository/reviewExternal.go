package review_external_repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

type ReviewExternalRepository struct {
}

func GetCountReviewsSc(scId int) int {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	var countReviews int

	conn.QueryRow(context.Background(), CountScReviews, scId).Scan(&countReviews)

	fmt.Println(countReviews)

	return countReviews
}
