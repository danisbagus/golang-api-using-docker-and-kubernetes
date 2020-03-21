package main

import (
	"fmt"
	"log"
	"os"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //for import mysql

	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", username, password, host, port, dbName)
	db, err := gorm.Open("mysql", dbURI)

	if err != nil {
		log.Fatalf("failed to connect to database:",err)
	}

	fmt.Println("success connect to database")

	defer db.Close()
}