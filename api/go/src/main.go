package main

import (
	"main/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/sample", services.GetAllSamples)
	router.GET("/sample/:id", services.GetSampleByID)
	router.POST("/sample", services.PostSample)
	router.PUT("/sample", services.PutSample)
	router.DELETE("/sample/:id", services.DeleteSample)

	router.Run(":8080")
}
