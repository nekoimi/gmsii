package router

import (
	"github.com/gin-gonic/gin"
	apiV1 "github.com/nekoimi/gmsii/apis/v1"
)

func InitApiRouter(e *gin.Engine)  {
	e.GET("/", apiV1.Welcome)
	v1 := e.Group("api/v1")
	{
		v1.POST("push", apiV1.OldPush)
		v1.POST("send/text", apiV1.Text)
		v1.POST("send/error", apiV1.Error)
	}
}
