package gorm

import "github.com/golang-module/carbon/v2"

type M struct {
	Id        uint            `json:"id" gorm:"primaryKey;comment:auto increment id" `
	CreatedAt carbon.DateTime `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt carbon.DateTime `json:"updatedAt" gorm:"comment:修改时间"`
	DeletedAt DeletedAt       `json:"deletedAt" gorm:"index:idx_deleted_at;comment:软删除时间"`
	UUID      string          `json:"uuid" gorm:"index:idx_uuid;comment:uuid"`
}
