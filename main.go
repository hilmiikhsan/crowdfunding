package main

import (
	"crowdfunding/handler"
	"crowdfunding/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:#21012123Op@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// repository
	userRepository := user.NewRepository(db)

	// service
	userService := user.NewService(userRepository)

	// handler
	userHandler := handler.NewUserHandler(userService)

	// route
	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/register-user", userHandler.RegisterUser)

	router.Run()
}
