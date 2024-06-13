package main

import (
	"github.com/gin-gonic/gin"
	"gin/unit-test-crud-user/entity"
	"gin/unit-test-crud-user/handler"
	"gin/unit-test-crud-user/repository/slice"
	"gin/unit-test-crud-user/router"
	"gin/unit-test-crud-user/service"
	"log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// setup service
	var mockUserDBInSlice []entity.User
	userRepo := slice.NewUserRepository(mockUserDBInSlice)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Routes
	router.SetupRouter(r, userHandler)

	// Run the server
	log.Println("Running server on port 8080")
	r.Run(":8080")
}
