package postgresql

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/ddd/entity"
	"github.com/gitHomesPhp/repaircat/ddd/valueobject"
	"github.com/jackc/pgx/v4"
	"os"
)

type UndergroundRepo struct {
	undergrounds            []*entity.Underground
	undergroundValueObjects []*valueobject.Underground
}

func GetRepo() *UndergroundRepo {
	return &UndergroundRepo{
		undergrounds:            []*entity.Underground{},
		undergroundValueObjects: []*valueobject.Underground{},
	}
}

func (undergroundRepo *UndergroundRepo) ListByCity(cityId int) (error, []*valueobject.Underground) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err, nil
	}

	rows, _ := conn.Query(context.Background(), GetUndergroundsByCity, cityId)

	for rows.Next() {
		underground := &valueobject.Underground{
			Id:    0,
			Label: "",
			Slug:  "",
		}

		rows.Scan(&underground.Label, &underground.Id, &underground.Slug)

		undergroundRepo.undergroundValueObjects = append(undergroundRepo.undergroundValueObjects, underground)
	}

	defer conn.Close(context.Background())

	return nil, undergroundRepo.undergroundValueObjects
}
