package sqlconnect

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var count = 1

func ConnectDB() (*pgx.Conn, error) {

	for {
		conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
		if err != nil {
			if count > 5 {
				log.Println("Couldn't connect to postgres DB")
				return nil, err
			}
			log.Println("retrying to connect ...")
			count++
		} else {
			return conn, nil
		}
	}

}
