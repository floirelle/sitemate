package main

import (
	"server/controller"
	"server/repository"
	"server/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	repository := repository.NewRepository()
	service := service.NewService(repository)
	controller := controller.NewController(service)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PATCH"},
	}))
	router.GET("/issues", controller.GetIssues)
	router.GET("/issues/:id", controller.GetIssue)
	router.PATCH("/issues/:id", controller.UpdateIssue)
	router.POST("/issues", controller.AddIssue)
	router.DELETE("/issues/:id", controller.DeleteIssue)

	router.Run()
}
