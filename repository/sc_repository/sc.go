package sc_repository

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/entity"
	"github.com/gitHomesPhp/repaircat/utils/pg"
)

func List(page int) ([]map[string]any, error) {
	const COUNT = 10
	to := COUNT * page
	from := (COUNT*page+1)%COUNT + COUNT*(page-1)

	rows, err := pg.Conn().Query(context.Background(), SelectScPage, from, to)
	fmt.Println(rows)
	if err != nil {
		return nil, err
	}

	var scList []map[string]any
	i := 0
	for rows.Next() {
		sc := entity.EmptySc()

		err := rows.Scan(sc.GetAttributes())
		if err != nil {
			return nil, err
		}

		scList = append(scList, sc.ToMap())
		i++
	}

	var next int
	var previous int

	err = pg.Conn().QueryRow(context.Background(), SelectId, to+1).Scan(&next)
	if err != nil {
		return nil, err
	}
	err = pg.Conn().QueryRow(context.Background(), SelectId, from-1).Scan(&previous)
	if err != nil {
		return nil, err
	}

	scList = append(scList, map[string]any{
		"next":     next,
		"previous": previous,
	})

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
