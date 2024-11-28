package models

import (
	"time"

	//"github.com/jinzhu/gorm"
	"github.com/kbelyako/go/TasksManagementApp/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Task struct {
	gorm.Model
	Id              int
	User_id         int
	Title           string `gorm:""json:"title"`
	Description     string `json:"description"`
	Due_date        time.Time
	Is_reminder_set time.Time
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Task{})
}

func (t *Task) CreateTask() *Task {
	//db.NewRecord(t)
	db.Create(&t)
	return t
}

func GetTaskByID(Id int64) (*Task, *gorm.DB) {
	var getTask Task
	db := db.Where("ID=?", Id).Find(&getTask)
	return &getTask, db
}

func DeleteTask(ID int64) Task {
	var task Task
	db.Where("ID=?", ID).Delete(task)
	return task
}

func GetAllTasks() []Task {
	var Tasks []Task
	db.Find(&Tasks)
	return Tasks
}
