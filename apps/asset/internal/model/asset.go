package model

import "gorm.io/gorm"

type Asset struct {
	gorm.Model
	UUID            string  `json:"uuid" gorm:"comment:资产UUID"`
	ProjectId       string  `json:"projectId" gorm:"comment:资产项目ID"`
	SonProjectId    string  `json:"sonProjectId" gorm:"comment:资产子项目ID"`
	PartProjectId   string  `json:"partProjectId" gorm:"comment:资产分部项目ID"`
	ProjectName     string  `json:"projectName" gorm:"comment:项目名称：冗余字段"`
	SonProjectName  string  `json:"sonProjectName" gorm:"comment:子项目名称：冗余字段"`
	PartProjectName string  `json:"partProjectName" gorm:"comment:分部项目名称：冗余字段"`
	Serial          string  `json:"serial" gorm:"comment:资产编号"`
	Name            string  `json:"name" gorm:"comment:名称"`
	Type            string  `json:"type" gorm:"comment:资产类别"`
	Kind            string  `json:"kind" gorm:"comment:资产属性"`
	SubDistrict     string  `json:"subDistrict" gorm:"comment:街道" `
	Brand           string  `json:"brand" gorm:"comment:品牌"`
	Specs           string  `json:"specs" gorm:"comment:型号"`
	Unit            string  `json:"unit" gorm:"comment:单位"`
	Params          string  `json:"params" gorm:"comment:型号规格'"`
	Value           float64 `json:"value" gorm:"comment:价值"`
	Address         string  `json:"address" gorm:"comment:地址"`
	Long            string  `json:"long" gorm:"comment:经度"`
	Lat             string  `json:"lat" gorm:"comment:纬度"`
	ImgUrl          string  `json:"imgUrl" gorm:"comment:资产图片"`
	SystemLoginUrl  string  `json:"systemLoginUrl" gorm:"comment:系统登录地址"`
	SystemLoginPwd  string  `json:"systemLoginPwd" gorm:"comment:系统登录密码"`
	IotNetSerial    string  `json:"iotNetSerial" gorm:"comment:物联网卡/宽带号"`
	NetStatus       string  `json:"netStatus" gorm:"comment:网络资费情况"`
	EmeterSerial    string  `json:"emeterSerial" gorm:"comment:电表号"`
	EmeterStatus    string  `json:"emeterStatus" gorm:"comment:电费使用情况"`
}
