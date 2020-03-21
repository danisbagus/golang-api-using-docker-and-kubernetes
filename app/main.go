package main

import (
	router "go-project/routers"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var status string

func main() {

	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")

	r := router.SetupRouters()
	r.Run(":" + port)
	
}