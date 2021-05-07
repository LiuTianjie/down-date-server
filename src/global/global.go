package global

import (
	"down-date-server/src/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	CONFIG config.Server
	VP     *viper.Viper
)
