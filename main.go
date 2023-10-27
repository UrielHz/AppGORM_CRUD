package main

import (
	"AppGoORMCRUD/models"
	"AppGoORMCRUD/routes"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:1234@tcp(localhost:3306)/exampledb?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.Todo{})
	r := routes.SetupRouter(db)
	r.Run(":8080")
}
