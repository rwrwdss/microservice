package main

import (
	"fmt"

	"github.com/rwrwdss/microservices-cache-app/services/user-service/internal/config"
	"github.com/rwrwdss/microservices-cache-app/services/user-service/internal/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	_, err = database.Connect(cfg)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}

	fmt.Println("User service is running...")
}
