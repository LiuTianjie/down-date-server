package router

import (
	"down-date-server/src/api"

	"github.com/gin-gonic/gin"
)

func InitBaseRoute(Router *gin.RouterGroup) {
	UserRouter := Router.Group("")
	{
		UserRouter.POST("register", api.Register)
		UserRouter.POST("login", api.Login)
	}
}

func IintAuthRoute(Router *gin.RouterGroup) {
	UserRouter := Router.Group("")
	{
		UserRouter.GET("search_all_user", api.SearchUser)
		UserRouter.GET("search_user", api.SearchUserByNickname)
	}
}
