/*
 * @Descripttion: initialize danmu
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-03 10:46:29
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-10 16:41:31
 */
package main

import (
	"down-date-server/src/global"
	"down-date-server/src/initialize"
)

func main() {
	global.VP = initialize.Viper()
	global.DB = initialize.ConnectDB()
	// global.DANMU = initialize.InitDanmu()
	if global.DB != nil {
		initialize.Gorm(global.DB)
		db, _ := global.DB.DB()
		defer db.Close()
	}
	router := initialize.Routers()
	router.Run(":8000")
}
