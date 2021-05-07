/*
 * @Descripttion: note
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-06 23:16:02
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-07 12:05:25
 */
package config

// 服务相关的配置，通过config.yaml来初始化
// 目前有JWT，数据库相关配置
type Server struct {
	JWT        JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	ServerInfo string `mapstructure:"server-info" json:"server-info" yaml:"server-info"`
	Mysql      Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
