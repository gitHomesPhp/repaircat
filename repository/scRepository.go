package repository

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/types"
	"github.com/gitHomesPhp/repaircat/utils/pg"
)

const sqlSelectList = "SELECT " +
	"sc.id, " +
	"COALESCE(name, ''), " +
	"COALESCE(description, ''), " +
	"COALESCE(phone, ''), " +
	"COALESCE(email, ''), " +
	"COALESCE(site, ''), " +
	"COALESCE(city, ''), " +
	"COALESCE(address, ''), " +
	"COALESCE(underground, '') " +
	"FROM sc LEFT JOIN location " +
	"ON location_id = location.id "

const insertSc = "insert into sc(name, description, phone, email, site, location_id) values($1, $2, $3, $4, $5, $6)"

func AddSc(sc *types.Sc) {
	_, err := pg.Conn().Exec(context.Background(), insertSc, sc.Name, sc.Description, sc.Phone, sc.Email, sc.Site, sc.LocationId)
	fmt.Println(err)
}

func GetScList2(pageNumber int) map[int]map[string]any {
	const COUNT = 10
	to := COUNT * pageNumber
	from := (COUNT*pageNumber+1)%COUNT + COUNT*(pageNumber-1)

	rows, err := pg.Conn().Query(context.Background(), sqlSelectList+"where sc.id >= $1 and sc.id <= $2", from, to)

	if err != nil {
		fmt.Println(err)
	}

	var scList = make(map[int]map[string]any)
	i := 0
	for rows.Next() {
		var id int
		sc := types.NewSc()
		location := types.NewLocation()

		err := rows.Scan(&id, &sc.Name, &sc.Description, &sc.Phone, &sc.Email, &sc.Site, &location.City, &location.Address, &location.Underground)
		if err != nil {
			fmt.Println(err)
		}

		sc.SetId(id)
		scList[i] = sc.ToMap(location.City, location.Address, location.Underground)
		i++
	}

	var next int
	var previous int

	pg.Conn().QueryRow(context.Background(), "SELECT id FROM sc where id = $1", to+1).Scan(&next)
	pg.Conn().QueryRow(context.Background(), "SELECT id FROM sc where id = $1", from-1).Scan(&previous)

	scList[i] = map[string]any{
		"next":     next,
		"previous": previous,
	}
	return scList
}
