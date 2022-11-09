package model

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	UUID            string `json:"uuid" gorm:"comment:项目UUID"`
	Name            string `json:"name" gorm:"unique;comment:项目名称"`
	ParentProjectId string `json:"parentProjectId" gorm:"comment:父级项目ID"`
	Level           int64  `json:"level" gorm:"comment:项目层级"`
}
