package repository

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/types"
	"github.com/gitHomesPhp/repaircat/utils/pg"
)

const sqlSelectList = "select id, name, description, phone, email, site from sc "

func AddSc(sc *types.Sc) {
	_, err := pg.Conn().Exec(context.Background(), "insert into sc(name, description) values($1, $2)", sc.Name, sc.Description)
	fmt.Println(err)
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
