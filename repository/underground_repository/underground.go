package underground_repository

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/entity"
	"github.com/jackc/pgx/v4"
	"os"
)

func GetUndergroundByCityId(cityId int) ([]map[string]any, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	rows, err := conn.Query(context.Background(), SelectUndergroundByCityId, cityId)

	if err != nil {
		return nil, err
	}

	var undergrounds []map[string]any

	for rows.Next() {
		underground := entity.EmptyUnderground()

		rows.Scan(underground.GetAttributes()...)

		undergrounds = append(undergrounds, underground.ToMap())
	}

	defer conn.Close(context.Background())

	return undergrounds, nil
}
