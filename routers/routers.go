package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rn-consider/compuBackend/service/info"
)

func RegisterRoutes(router *gin.RouterGroup) {
	infoRouter := router.Group("/infos")

	// 注册创建信息的路由
	infoRouter.POST("/", info.CreateInfo)

	// 获取所有信息的路由
	infoRouter.GET("/", info.GetAllInfo)

	// 注册更新信息的路由
	infoRouter.PUT("/:id", info.UpdateInfo)

	// 注册删除信息的路由
	infoRouter.DELETE("/:id", info.DeleteInfo)

	// 注册获取信息的路由
	infoRouter.GET("/:id", info.GetInfo)

	// 注册更多路由...
}
