/*
 * @Descripttion: your project
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-06 15:16:05
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-10 14:39:05
 */
package global

import (
	"down-date-server/src/config"
	"down-date-server/src/danmu"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	CONFIG config.Server
	VP     *viper.Viper
	DANMU  *danmu.Danmu
)
