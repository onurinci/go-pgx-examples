package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:112233@213.238.179.17:5432/AygunEbsShift")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "select * from \"public\".\"Employees\" order by \"fullname\" asc limit 10") // \"public\".\"Employees\"
	if err != nil {
		log.Fatal("error while executing query")
	}

	// iterate through the rows
	log.Println(time.Now())
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}

		id := values[0].(int32)
		fullname := values[1].(string)

		log.Println("[id:", id, ", fullname:", fullname, "]")
	}
	log.Println(time.Now())
}
