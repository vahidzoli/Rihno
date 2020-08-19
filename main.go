package main

import (
	"Rihno/config"
	"Rihno/entity"
	"Rihno/routes"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.App.Init()
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	//config.DB.LogMode(true)
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close()

	config.DB.AutoMigrate(&entity.User{}, &entity.Project{}, &entity.Plan{})
	r := routes.SetupRouter()
	r.Run()
}
