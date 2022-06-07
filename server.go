package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/controllers"
	"github.com/jackc/pgx/v4"
	"os"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:secret@localhost:5432/repaircat")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name string
	err = conn.QueryRow(context.Background(), "select name from sc where sc_id=$1", 1).Scan(&name)

	println(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	r := gin.Default()
	r.GET("/ping", controllers.PongController)
	r.GET("/sc/:id", controllers.GetSC)
	r.Run()
}
