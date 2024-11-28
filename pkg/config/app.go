package config

import (
	"fmt"

	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/glebarez/sqlite"
	_ "github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	//conStringPG := "User=postgres;Password=54321!;Host=localhost;Port=5432;Database=myDataBase;Pooling=true;MinPoolSize=0;MaxPoolSize=100;ConnectionLifetime=0;"
	conStringSL := "Data Source=D:\\Dev\\Go\\TasksManagementApp\\db\\TasksDB.db;Version=3;"
	//d, err := gorm.Open("postgres", conString)
	//d, err := gorm.Open("sqlite", conStringSL)
	//d, err := gorm.Open(sqlite.Open, conStringSL, &gorm.Config{})
	d, err := gorm.Open(sqlite.Open(conStringSL), &gorm.Config{})
	if err != nil {
		fmt.Println("error while Starting DB")
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
