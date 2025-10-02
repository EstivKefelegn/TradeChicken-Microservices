package main

import (
	"context"
	"log"

	"github.com/product-service/internal/repository/sqlconnect"
)

func main() {
	log.Println("Product service is running")
	connction, err := sqlconnect.ConnectDB()
	if err != nil {
		log.Println("Couldn't connect to postgres database")
		return
	}
	ctx := context.Background()
	defer connction.Close(ctx)
	
	log.Println("Postgres DB is connected successfuly")
}
