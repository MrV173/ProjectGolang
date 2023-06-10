package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func DatabaseConnect() {

	databaseUrl := "postgres://postgres:admin@localhost:5432/Dumbways-Project-App"

	var err error
	Conn, err = pgx.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v/n", err)
		os.Exit(1)
	}

	fmt.Println("Succesfully connected to database!")
}
