package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

func main() {
	dbCon := db.NewDB()
	defer fmt.Println("Success")
	defer db.CloseDB(dbCon)
	dbCon.AutoMigrate(&model.User{}, &model.Task{})
}
