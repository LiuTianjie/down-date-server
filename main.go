package main

import (
	"down-date-server/src/global"
	"down-date-server/src/initialize"
)

func main() {
	global.VP = initialize.Viper()
	global.DB = initialize.ConnectDB()
	if global.DB != nil {
		initialize.Gorm(global.DB)
		db, _ := global.DB.DB()
		defer db.Close()
	}
	router := initialize.Routers()
	router.Run(":8000")
}
