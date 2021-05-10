/*
 * @Descripttion: your project
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-03 23:33:11
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-08 09:41:49
 */
package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID        uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`                                                    // 用户UUID
	Username    string    `json:"username" gorm:"comment:用户登录名"`                                                 // 用户登录名
	Password    string    `json:"password"  gorm:"comment:用户登录密码"`                                               // 用户登录密码
	Nickname    string    `json:"nickname" gorm:"default:系统用户;comment:用户昵称"`                                     // 用户昵称"
	HeaderImg   string    `json:"headerimg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"` // 用户头像
	AuthorityId string    `json:"authorityid" gorm:"default:888;comment:用户角色ID"`                                 // 用户角色ID
}
