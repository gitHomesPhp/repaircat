package sc_repository

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/entity"
	"github.com/gitHomesPhp/repaircat/repository/underground_repository"
	"github.com/jackc/pgx/v4"
	"os"
)

func Flush(sc *entity.Sc) error {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	_, err = conn.Exec(context.Background(), InsertSc,
		sc.GetName(),
		sc.GetDescription(),
		sc.GetPhone(),
		sc.GetEmail(),
		sc.GetSite(),
		sc.GetLocation().GetId(),
	)

	if err != nil {
		return err
	}

	defer conn.Close(context.Background())

	return nil
}

func Find(id int) (map[string]any, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	sc := entity.EmptySc()
	location := entity.EmptyLocation()

	attrs := append(sc.GetAttributes2(), location.GetAttributes()...)

	err = conn.QueryRow(context.Background(), SelectScById, id).Scan(attrs...)

	underground_repository.GetUndergroundsOfLocations([]*entity.Location{location})

	if err != nil {
		return nil, err
	}

	sc.AddLocation(location)

	defer conn.Close(context.Background())

	return sc.ToMap(), nil
}

func deleteElemSlice(s []any, n int) []any {
	s = append(s[:n], s[n+1:]...)

	return s
}
