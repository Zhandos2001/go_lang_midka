package main

import (
	c "midterm-project/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	controller:= c.GetAppController()
	router.GET("/", controller.HandleBaseRoute)
	router.GET("/store/:key", controller.GetValue)
	router.PUT("/store/update", controller.SetValue)
	router.Run()
}