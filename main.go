package main

import (
	"projects/bookcrud/models"
	"projects/bookcrud/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)
	r.Run()
}
