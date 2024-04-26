package handler

import (
	"github.com/gin-gonic/gin"
	G "github.com/kmou424/yukabot/global"
)

var baseRouter *gin.RouterGroup

func Register() {
	baseRouter = G.Router.Group("/api")
	v1()
}

func v1() {
	v1Router := baseRouter.Group("/v1")

	alarmGroup := v1Router.Group("/alarm")
	{
		alarmGroup.POST("/text", alarmText)
	}
}
