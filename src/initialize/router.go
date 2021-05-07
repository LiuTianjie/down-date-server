package initialize

import (
	"down-date-server/src/middle"
	"down-date-server/src/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	PublicGroup := Router.Group("base")
	{
		router.InitBaseRoute(PublicGroup)
	}
	PrivateGroup := Router.Group("v1")
	PrivateGroup.Use(middle.JWTAuth())
	{
		router.IintAuthRoute(PrivateGroup)
	}
	return Router
}
