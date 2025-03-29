package main

import (
	"fmt"
	"learning-go-rest-api/db"
	"learning-go-rest-api/model"
	"log"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully migrate")
	defer db.CloseDB(dbConn)
	if err := dbConn.AutoMigrate(&model.User{}, &model.Task{}); err != nil {
		log.Fatalln(err)
	}
}
