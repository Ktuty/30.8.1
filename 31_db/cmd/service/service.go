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
	//! Добавление
	// task, err := db.NewTask(postgres.Task{
	// 	Title:   "New Task",
  //   Content: "This is a new task.",
	// })

	// if err != nil {
  //   log.Fatal(err)
  // }
	// fmt.Println("New task ID:", task)

	//! Редактирование
	// fmt.Println(db.EditTask(postgres.Task{
	// 	ID: 8,
	//   Title: "Edit Task", 
	// 	Content: "XXX",
	// }))

	//! Удаление
	//db.DeleteTask(postgres.Task{ID: 8})

	//! Получение всех задач
	tasks, err := db.Tasks(0, 0)
	if err != nil {
    log.Fatal(err)
  }

	fmt.Println(tasks)


	//! Получение задач по лейблу (1)
	// tasks, err := db.TasksByLabel(postgres.Label{
	// 	Name: "Маша",
	// })
	// if err != nil {
  //   log.Fatal(err)
  // }

	// fmt.Println("Маша", tasks)

	//! Получение задач по лейблу (2)
	// tasks, err = db.TasksByLabel(postgres.Label{
	// 	Name: "Вася",
	// })
	// if err != nil {
  //   log.Fatal(err)
  // }

	// fmt.Println("Вася", tasks)
}