package main

import (
	"fmt"
	"log"
	"modules/pkg/storage"
	"modules/pkg/storage/postgres"
)

var db storage.Interface

func main() {
	var err error

	// pwd := os.Getenv("pwdDB")
	// if pwd == ""{
	// 	os.Exit(1)
	// }

	connstr := "postgres://postgres:Taran@12092004@localhost/tasks?sslmode=disable"

	db, err = postgres.NewStorage(connstr)
	if err != nil {
    log.Fatal(err)
  }

	tasks, err := db.Tasks(0, 0)
	if err != nil {
    log.Fatal(err)
  }

	fmt.Println(tasks)
}