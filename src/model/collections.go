/*
 * @Descripttion: collections
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-07 14:31:18
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-08 09:42:43
 */
package model

import "gorm.io/gorm"

type Collections struct {
	gorm.Model
	Username   string
	Kind       string
	University string
	Profession string
}
