package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/kbelyako/go/TasksManagementApp/pkg/models"
	"github.com/kbelyako/go/TasksManagementApp/pkg/utils"
)

var NewTask models.Task

func GetTask(w http.ResponseWriter, r *http.Request) {
	NewTasks := models.GetAllTasks()
	res, _ := json.Marshal(NewTasks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["taskId"]
	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	taskDetails, _ := models.GetTaskByID(ID)
	res, _ := json.Marshal(taskDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	CreateTask := &models.Task{}
	utils.ParseBody(r, CreateTask)
	t := CreateTask.CreateTask()
	res, _ := json.Marshal(t)
	//w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	taskId := vars["taskId"]
	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	task := models.DeleteTask(ID)
	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	var updateTask = &models.Task{}
	utils.ParseBody(r, updateTask)
	vars := mux.Vars(r)
	taskId := vars["taskId"]
	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	taskDetails, db := models.GetTaskByID(ID)
	if updateTask.Title != "" {
		taskDetails.Title = updateTask.Title
	}
	if updateTask.Description != "" {
		taskDetails.Description = updateTask.Description
	}
	if updateTask.Due_date.IsZero() != true {
		taskDetails.Due_date = updateTask.Due_date
	}

	db.Save(&taskDetails)

	//task := models.DeleteTask(ID)
	res, _ := json.Marshal(taskDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
