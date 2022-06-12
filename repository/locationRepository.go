package repository

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/types"
	"github.com/gitHomesPhp/repaircat/utils/pg"
)

const insertLocation = "insert into location(city, address, underground) values($1, $2, $3)"

func AddLocation(location *types.Location) {
	_, err := pg.Conn().Exec(context.Background(), insertLocation, location.City, location.Address, location.Underground)
	fmt.Println(err)
}

func getLocationId() {

}
