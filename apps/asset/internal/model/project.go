package model

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	UUID            string `json:"uuid" gorm:"comment:project's id"`
	Name            string `json:"name" gorm:"unique;comment:project's name"`
	ParentProjectId string `json:"parentProjectId" gorm:"comment:parent project id"`
}
