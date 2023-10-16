package main

import (
	"fmt"
	"go_echo_api/db"
	"go_echo_api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated...")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.Post{})
}
