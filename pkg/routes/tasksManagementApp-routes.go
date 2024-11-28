package routes

import (
	"github.com/gorilla/mux"
	"github.com/kbelyako/go/TasksManagementApp/pkg/controllers"
)

var RegisterTasksManagementApp = func(router *mux.Router) {
	router.HandleFunc("/task/", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/task/", controllers.GetTask).Methods("GET")
	router.HandleFunc("/task/{taskId}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/task/{taskId}", controllers.DeleteTask).Methods("DELETE")
}
