package city_repository

import (
	"context"
	"fmt"
	"github.com/gitHomesPhp/repaircat/entity"
	"github.com/jackc/pgx/v4"
	"os"
)

func GetAll() {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:secret@postgres:5432/repaircat")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

}

func Flush(city *entity.City) {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:secret@postgres:5432/repaircat")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	_, err = conn.Exec(context.Background(), InsertCity,
		city.GetLabel(),
		city.GetCode(),
	)

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close(context.Background())
}
