package main

import (
	"database/sql"
	"gomod/service"
	"gomod/storage"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	userStore := &storage.UserStorage{UserStore: db}
	userService := service.UserService{UserServ: userStore}
	mux.HandleFunc("/create_user", userService.CreateUserHandler)
	mux.HandleFunc("/authorize_user", userService.UserAuthorizationHandler)

	taskStore := &storage.TaskStorage{TaskStore: db}
	taskService := service.TaskService{TaskServ: taskStore}
	mux.HandleFunc("/update", taskService.EditNameHandler)
	mux.HandleFunc("/update", taskService.EditLastNameHandler)
	mux.HandleFunc("/deliete", taskService.EditSexHandler)
	mux.HandleFunc("/get_all", taskService.EditFamilyStatusHandler)
	mux.HandleFunc("/get_some/", taskService.EditCityHandler)
	mux.HandleFunc("/get_one/", taskService.AddUnclesHandler)
	mux.HandleFunc("/update_status", taskService.DeleteUnclesHandler)
	mux.HandleFunc("/update_status", taskService.WriteMessageHandler)
	mux.HandleFunc("/update_status", taskService.UpdateMessageHandler)
	mux.HandleFunc("/update_status", taskService.DeleteMessageHandler)

	http.ListenAndServe("localhost:8080", mux)
}
