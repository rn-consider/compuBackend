package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/rn-consider/compuBackend/config"
	"github.com/rn-consider/compuBackend/routers"
	"github.com/rn-consider/compuBackend/routers/middleware"
	"github.com/rn-consider/compuBackend/utils"
	"net/http"
)

func HealthCheck(g *gin.Context) {
	g.JSON(http.StatusOK, "ok...")
}

func Routers() *gin.Engine {

	if err := utils.Translator("zh"); err != nil {
		config.GVA_LOG.Error(err.Error())
		return nil
	}

	Router := gin.Default()
	//gin.SetMode(gin.DebugMode)

	Router.Use(middleware.Recovery())
	Router.Use(middleware.Logger())
	Router.Use(middleware.Cors())
	HealthGroup := Router.Group("")
	{
		// 健康监测
		HealthGroup.GET("/health", HealthCheck)
	}
	// 注册路由
	routerGroup := Router.Group("")     // 创建一个路由组
	routers.RegisterRoutes(routerGroup) // 注册路由到路由组
	return Router
}
