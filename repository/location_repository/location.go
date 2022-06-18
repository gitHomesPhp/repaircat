package location_repository

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/entity"
	"github.com/jackc/pgx/v4"
	"os"
)

func Flush(location *entity.Location) (error, *entity.Location) {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:secret@localhost:5431/repaircat")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	_, err = conn.Exec(context.Background(), InsertLocation,
		location.GetCity(),
		location.GetAddress(),
		location.GetUnderGround(),
	)

	if err != nil {
		return err, nil
	}

	var id int

	err = conn.QueryRow(context.Background(), GetIdByValues, location.GetCity(), location.GetAddress(), location.GetUnderGround()).Scan(&id)

	if err != nil {
		return err, nil
	}

	location.SetId(id)

	defer conn.Close(context.Background())

	return nil, location
}
