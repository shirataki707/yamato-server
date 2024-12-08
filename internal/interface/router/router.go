package router

import (
	"yamato/internal/interface/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(mountainHandler *handler.MountainHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/mountains/:id", mountainHandler.GetMountain)
		api.GET("/mountains", mountainHandler.GetAllMountains)
	}

	return r
}
