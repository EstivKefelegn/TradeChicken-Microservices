package sqlconnect

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

var counts int32
var retryDelay = 3 * time.Second

func ConnectDB() (*pgx.Conn, error) {
	for {
		connection, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
		if err != nil {
			if counts > 5 {
				log.Println("Couldn't cpnnect to postgres")
				return nil, err
			}
			log.Println("Postgres not ready yet")
			counts++
		} else {
			log.Println("Connected to potgres")
			return connection, nil
		}
		time.Sleep(retryDelay)

	}

}
