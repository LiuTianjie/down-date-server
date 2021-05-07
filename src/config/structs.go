/*
 * @Descripttion: Add mysql struct
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-06 23:15:06
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-07 12:03:30
 */
package config

// 定义JWT的数据结构
type JWT struct {
	SignKey     string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`    // jwt签名
	ExpiresTime int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"` // 过期时间
	BufferTime  int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`    // 缓冲时间
}

// 定义mysql的结构
type Mysql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"` // 服务器地址:端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`     // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"` // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	LogZap       string `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
}
