package main

import (
	"context"
	"log"

	"github.com/order-service/internal/repository/sqlconnect"
)

func main() {
	log.Println("Order service is strting")
	conection, err := sqlconnect.ConnectDB()
	if err != nil {
		log.Println("Couldn't connect to db")
		return 
	}

	defer conection.Close(context.Background())

	log.Println("DB connected to postgres successfully")
}