package location_repository

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/entity"
	"github.com/jackc/pgx/v4"
	"os"
)

func Flush(location *entity.Location) (error, *entity.Location) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	_, err = conn.Exec(context.Background(), InsertLocation,
		location.GetCity().GetId(),
		location.GetAddress(),
		location.GetUnderGround().GetId(),
	)

	if err != nil {
		fmt.Println(err)
		return err, nil
	}

	var id int

	if location.GetUnderGround().GetId() == 0 {
		err = conn.QueryRow(context.Background(), SelectIdByValuesAndNullUnderground,
			location.GetCity().GetId(),
			location.GetAddress(),
		).Scan(&id)
	} else {
		err = conn.QueryRow(context.Background(), GetIdByValues,
			location.GetCity().GetId(),
			location.GetAddress(),
			location.GetUnderGround().GetId(),
		).Scan(&id)
	}

	if err != nil {
		fmt.Println(err)
		return err, nil
	}

	location.SetId(id)

	defer conn.Close(context.Background())

	return nil, location
}
