package main

import (
	"library-api/config"
	"library-api/models"
	"library-api/routes"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(mysql.Open(config.GetDSN()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.Book{}, &models.Author{}, &models.Customer{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	r := gin.Default()

	routes.RegisterBookRoutes(r, db)
	routes.RegisterAuthorRoutes(r, db)
	routes.RegisterCustomerRoutes(r, db)

	r.Run()
}
