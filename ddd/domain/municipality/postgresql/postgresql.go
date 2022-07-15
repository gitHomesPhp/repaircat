package postgresql

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/ddd/valueobject"
	"github.com/jackc/pgx/v4"
	"os"
)

type MunicipalityRepo struct {
	municipalityValueObjects []*valueobject.Municipality
}

func GetRepo() *MunicipalityRepo {
	return &MunicipalityRepo{
		municipalityValueObjects: []*valueobject.Municipality{},
	}
}

func (municipalityRepo *MunicipalityRepo) ListByCity(cityId int) (error, []*valueobject.Municipality) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err, nil
	}

	rows, _ := conn.Query(context.Background(), GetMunicipalitiesByCityId, cityId)

	for rows.Next() {
		municipality := &valueobject.Municipality{
			Id:    0,
			Label: "",
			Slug:  "",
		}

		rows.Scan(&municipality.Label, &municipality.Id, &municipality.Slug)

		municipalityRepo.municipalityValueObjects = append(municipalityRepo.municipalityValueObjects, municipality)
	}

	defer conn.Close(context.Background())

	return nil, municipalityRepo.municipalityValueObjects
}
