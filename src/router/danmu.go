/*
 * @Descripttion:Add danmu route
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-10 15:24:01
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-10 15:28:56
 */
package router

import (
	"down-date-server/src/api"

	"github.com/gin-gonic/gin"
)

func InitDanmuRoute(Router *gin.RouterGroup) {
	DanmuRouter := Router.Group("")
	{
		DanmuRouter.GET("danmu", api.DanmuHandler)
	}
}
