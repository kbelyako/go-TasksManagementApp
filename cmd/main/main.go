package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kbelyako/go/TasksManagementApp/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterTasksManagementApp(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhos:9010", r))
}
