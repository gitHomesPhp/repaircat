package postgresql

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/ddd/aggregate"
	"github.com/gitHomesPhp/repaircat/ddd/entity"
	"github.com/gitHomesPhp/repaircat/ddd/valueobject"
	"github.com/jackc/pgx/v4"
	"os"
)

type ScCardExtensionRepo struct {
	scCardExtension []*aggregate.ScCardExtension
}

func GetRepo() *ScCardExtensionRepo {
	return &ScCardExtensionRepo{
		scCardExtension: []*aggregate.ScCardExtension{},
	}
}

func (scCardExtensionRepo *ScCardExtensionRepo) Get(id int) (error, *aggregate.ScCardExtension) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err, nil
	}
	scCardExtension := aggregate.NewScCardExtension()

	row := conn.QueryRow(context.Background(), SelectScCardExtensionById, id)

	var locationId int

	row.Scan(
		&scCardExtension.Sc.Id,
		&scCardExtension.Sc.Name,
		&scCardExtension.Sc.Description,
		&scCardExtension.Sc.Phone,
		&scCardExtension.Sc.Email,
		&scCardExtension.Sc.Site,
		&scCardExtension.Sc.Location.City.Label,
		&scCardExtension.Sc.Location.Address,
		&scCardExtension.Sc.Location.Latitude,
		&scCardExtension.Sc.Location.Longitude,
		&locationId,
	)

	rows, _ := conn.Query(context.Background(), GetUndergrounds, locationId)

	for rows.Next() {
		underground := &entity.Underground{Label: ""}

		rows.Scan(&underground.Label)

		scCardExtension.Sc.Location.Undergrounds = append(scCardExtension.Sc.Location.Undergrounds, underground)
	}

	rows, _ = conn.Query(context.Background(), GetMunicipalities, locationId)

	for rows.Next() {
		municipality := &valueobject.Municipality{
			Id:    0,
			Label: "",
			Slug:  "",
		}

		rows.Scan(&municipality.Label)

		scCardExtension.Sc.Location.Municipalities = append(scCardExtension.Sc.Location.Municipalities, municipality)
	}

	reviewInfo := &valueobject.ReviewInfo{
		Count:  0,
		Rating: 0,
	}
	rows.Scan(
		&reviewInfo.Count,
		&reviewInfo.Rating,
	)

	conn.QueryRow(context.Background(), GetReviewInfo, id).Scan(&reviewInfo.Count, &reviewInfo.Rating)

	scCardExtension.ReviewInfo = reviewInfo

	defer conn.Close(context.Background())

	return nil, scCardExtension
}
