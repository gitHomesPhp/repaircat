package city_repository

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/entity"
	"github.com/jackc/pgx/v4"
	"os"
)

func GetAllCities() ([]map[string]any, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	rows, err := conn.Query(context.Background(), SelectCities)
	if err != nil {
		return nil, err
	}

	var cities []map[string]any

	for rows.Next() {
		city := entity.EmptyCity()

		rows.Scan(city.GetAttributes()...)
		cities = append(cities, city.ToMap())
	}

	defer conn.Close(context.Background())

	return cities, nil
}

func Flush(city *entity.City) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	_, err = conn.Exec(context.Background(), InsertCity,
		city.GetLabel(),
		city.GetCode(),
	)

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close(context.Background())
}
