package main

import (
	"MaintenanceSystem/handler"
	"MaintenanceSystem/repository"
	"MaintenanceSystem/service"

	"github.com/gin-gonic/gin"
)

func main() {

	db, _ := repository.InitDB()

	dbs := service.InitServiceDB(db)

	r := gin.Default()

	userHandler := handler.NewUserHandler(dbs)

	r.POST("/login", userHandler.Login)
	auth := r.Group("/api", middleware.AuthMiddleware())

    auth.GET("/profile", userHandler.Profile)

	r.Run(":8080")

}
