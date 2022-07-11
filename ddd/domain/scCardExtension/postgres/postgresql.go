package postgres

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/ddd/aggregate"
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

	row.Scan(&scCardExtension.Sc.Name,
		&scCardExtension.Sc.Description,
		&scCardExtension.Sc.Phone,
		&scCardExtension.Sc.Email,
		&scCardExtension.Sc.Site,
		&scCardExtension.
	)

	defer conn.Close(context.Background())

	return nil, scCardExtensionRepo.scCardExtension[0]
}
