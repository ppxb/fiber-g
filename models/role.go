package models

import "fiber-g/pkg/gorm"

const (
	RoleStatusDisabled uint = 0
	RoleStatusEnable   uint = 1
)

type Role struct {
	gorm.M
	Name   string `json:"name" gorm:"comment:角色名称"`
	Status *uint  `json:"status" gorm:"type:tinyint(1);default:1;comment:部门状态(0-禁用 1-启用)"`
	Sort   *uint  `json:"sort" gorm:"default:1;comment:该字段必须≥0，值越大则权限越大，超级管理员为0"`
	Menus  []uint `json:"menus" gorm:"-"`
	Users  []User `json:"users" gorm:"foreignKey:RoleId"`
}
