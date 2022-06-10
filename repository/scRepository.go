package repository

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/types"
	"github.com/gitHomesPhp/repaircat/utils/pg"
)

const sqlSelectList = "select id, name, description, phone, email, site from sc "
const insertSc = "insert into sc(name, description, phone, email, site) values($1, $2, $3, $4, $5)"

func AddSc(sc *types.Sc) {
	_, err := pg.Conn().Exec(context.Background(), insertSc, sc.Name, sc.Description, sc.Phone, sc.Email, sc.Site)
	fmt.Println(err)
}

func GetScList2(pageNumber int) map[int]map[string]any {
	const COUNT = 10
	to := COUNT * pageNumber
	from := (COUNT*pageNumber+1)%COUNT + COUNT*(pageNumber-1)

	rows, err := pg.Conn().Query(context.Background(), sqlSelectList+"where id >= $1 and id <= $2", from, to)

	if err != nil {
		fmt.Println(err)
	}

	var scList = make(map[int]map[string]any)
	i := 0
	for rows.Next() {
		var id int
		sc := types.NewSc()

		err := rows.Scan(&id, &sc.Name, &sc.Description, &sc.Phone, &sc.Email, &sc.Site)
		if err != nil {
			fmt.Println(err)
			fmt.Println(err)
			fmt.Println(err)
		}

		sc.SetId(id)
		scList[i] = sc.ToMap()
		i++
	}

	return scList
}

func GetScList(pageNumber int) [10]*types.Sc {
	const COUNT = 10
	to := COUNT * pageNumber
	from := (COUNT*pageNumber+1)%COUNT + COUNT*(pageNumber-1)

	rows, err := pg.Conn().Query(context.Background(), sqlSelectList+"where id >= $1 and id <= $2", from, to)

	if err != nil {
		fmt.Println(err)
	}

	var scList [10]*types.Sc
	i := 0
	for rows.Next() {
		var id int
		sc := types.NewSc()

		err := rows.Scan(&id, &sc.Name, &sc.Description, &sc.Phone, &sc.Email, &sc.Site)
		if err != nil {
			fmt.Println(err)
		}

		sc.SetId(id)

		scList[i] = sc
		i++
	}

	return scList
}
