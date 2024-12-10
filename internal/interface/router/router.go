package router

import (
	"net/http"
	"yamato/internal/interface/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(mountainHandler *handler.MountainHandler) *gin.Engine {
	r := gin.Default()

	r.Static("/images", "./public/images")

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "HelloWorld")
	})

	api := r.Group("/api")
	{
		api.GET("/mountains/:id", mountainHandler.GetMountain)
		api.GET("/mountains", mountainHandler.GetAllMountains)
	}

	return r
}
