package main

import (
	"github.com/gin-gonic/gin"
	"github.com/supra08/hypertrace-demo-events-app/controllers"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1/events")
	{
		v1.POST("/", controllers.CreateEvent)
		v1.GET("/", controllers.FetchAllEvents)
		v1.GET("/:id", controllers.FetchEvent)
		v1.PUT("/:id", controllers.UpdateEvent)
		v1.DELETE("/:id", controllers.DeleteEvent)
	}
	router.Run()
}
