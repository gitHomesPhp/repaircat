package postgresql

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/ddd/aggregate"
	"github.com/jackc/pgx/v4"
	"os"
)

type ReviewForScCardExtensionRepository struct {
	Reviews []*aggregate.ReviewForScCardExtension
}

func GetRepo() *ReviewForScCardExtensionRepository {
	return &ReviewForScCardExtensionRepository{
		Reviews: []*aggregate.ReviewForScCardExtension{},
	}
}

func (repo *ReviewForScCardExtensionRepository) ListBySc(page int, scId int) (error, []*aggregate.ReviewForScCardExtension, map[string]bool) {
	const COUNT = 10

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err, nil, nil
	}

	rows, _ := conn.Query(context.Background(), SelectScReviewsExternalByCursor, scId, COUNT, COUNT*page-COUNT)

	for rows.Next() {
		review := aggregate.NewReviewForScCardExtension()
		var visitorId int

		rows.Scan(&review.Review.Id, &review.Review.Text, &review.Review.Rating, &visitorId)

		repo.Reviews = append(repo.Reviews, review)
	}

	next := true
	previous := true

	defer conn.Close(context.Background())

	return nil, repo.Reviews, map[string]bool{
		"previous": previous,
		"next":     next,
	}
}
