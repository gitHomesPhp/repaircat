package postgresql

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/ddd/entity"
	"github.com/jackc/pgx/v4"
	"os"
)

type CityRepo struct {
	Cities []*entity.City
}

func GetRepo() *CityRepo {
	return &CityRepo{Cities: []*entity.City{}}
}

func (cityRepo *CityRepo) List() (error, []*entity.City) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err, nil
	}

	rows, _ := conn.Query(context.Background(), "SELECT code, label FROM city")

	for rows.Next() {
		city := &entity.City{
			Code:  "",
			Label: "",
		}

		rows.Scan(&city.Code, &city.Label)

		cityRepo.Cities = append(cityRepo.Cities, city)
	}

	defer conn.Close(context.Background())

	return nil, cityRepo.Cities
}
