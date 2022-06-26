package sc_repository

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/entity"
	"github.com/gitHomesPhp/repaircat/repository/underground_repository"
	"github.com/jackc/pgx/v4"
	"os"
)

func List(page int) ([]map[string]any, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	const COUNT = 10
	to := COUNT * page
	from := (COUNT*page+1)%COUNT + COUNT*(page-1)

	rows, _ := conn.Query(context.Background(), SelectScWithLocationPageCursor, COUNT, COUNT*page-COUNT)

	var scList []map[string]any
	var scSlice []*entity.Sc
	var locations []*entity.Location

	for rows.Next() {
		sc := entity.EmptySc()
		location := entity.EmptyLocation()

		locAttr := deleteElemSlice(deleteElemSlice(location.GetAttributes(), 4), 3)

		attrs := append(sc.GetAttributes2(), locAttr...)
		err := rows.Scan(attrs...)

		if location.GetId() != 0 {
			sc.AddLocation(location)
		}

		if err != nil {
			return nil, err
		}

		locations = append(locations, location)
		scSlice = append(scSlice, sc)
	}

	underground_repository.GetUndergroundsOfLocations(locations)

	for i := 0; i < len(scSlice); i++ {
		scList = append(scList, scSlice[i].ToMap())
	}

	var next int
	var previous int

	_ = conn.QueryRow(context.Background(), SelectId, to+1).Scan(&next)

	_ = conn.QueryRow(context.Background(), SelectId, from-1).Scan(&previous)

	scList = append(scList, map[string]any{
		"next":     next,
		"previous": previous,
	})

	defer conn.Close(context.Background())

	return scList, nil
}

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
