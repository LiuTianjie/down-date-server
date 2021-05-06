package global

import (
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	JWT JWT_CONFIG
)

type JWT_CONFIG struct {
	SignKey     string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`    // jwt签名
	ExpiresTime int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"` // 过期时间
	BufferTime  int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`    // 缓冲时间
}
