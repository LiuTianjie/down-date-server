package initialize

import (
	"down-date-server/src/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	PublicGroup := Router.Group("base")
	{
		router.InitBaseRoute(PublicGroup)
	}
	return Router
}
