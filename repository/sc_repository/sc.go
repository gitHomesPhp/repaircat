package sc_repository

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/entity"
	"github.com/gitHomesPhp/repaircat/utils/pg"
	"github.com/jackc/pgx/v4"
	"os"
)

func List(page int) ([]map[string]any, error) {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:secret@localhost:5431/repaircat")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	const COUNT = 10
	to := COUNT * page
	from := (COUNT*page+1)%COUNT + COUNT*(page-1)

	rows, _ := conn.Query(context.Background(), SelectScPage, from, to)

	var scList []map[string]any

	for rows.Next() {
		sc := entity.EmptySc()

		err := rows.Scan(sc.GetAttributes())
		if err != nil {
			return nil, err
		}

		scList = append(scList, sc.ToMap())
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
	_, err := pg.Conn().Exec(context.Background(), InsertSc,
		sc.GetName(),
		sc.GetDescription(),
		sc.GetPhone(),
		sc.GetEmail(),
		sc.GetSite(),
	)

	if err != nil {
		return err
	}

	return nil
}
