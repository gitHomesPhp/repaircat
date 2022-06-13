package repository

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/types"
	"github.com/gitHomesPhp/repaircat/utils/pg"
)

const insertLocation = "insert into location(city, address, underground) values($1, $2, $3)"
const sqlLocationId = "select id from location where city = $1 and address = $2 and underground = $3"

func AddLocation(location *types.Location) {
	_, err := pg.Conn().Exec(context.Background(), insertLocation, location.City, location.Address, location.Underground)
	fmt.Println(err)
}

func GetLocationId(location *types.Location) int {
	var id int

	err := pg.Conn().QueryRow(context.Background(), sqlLocationId, location.City, location.Address, location.Underground).Scan(&id)

	if err != nil {
		fmt.Println(err)
	}

	return id
}
