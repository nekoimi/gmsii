package v1

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Welcome(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": 0,
		"data": gin.H{
			"time": time.Now().Format("2006-01-02 15:04:05"),
		},
		"message": "golang大法好",
	})
}
