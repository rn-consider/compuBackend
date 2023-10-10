package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rn-consider/compuBackend/service/organization_info"
)

func RegisterRoutes(router *gin.RouterGroup) {
	organizationRouter := router.Group("/organizations")

	organizationRouter.POST("/", organization_info.CreateOrganization())
	organizationRouter.PUT("/:id", updateOrganization)
	organizationRouter.DELETE("/:id", deleteOrganization)
	organizationRouter.GET("/:id", getOrganization)
	// 注册更多路由...
}
