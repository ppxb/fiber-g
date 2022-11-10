package models

import (
	"fiber-g/pkg/gorm"
	"github.com/golang-module/carbon/v2"
)

const (
	UserStatusDisabled uint = 0
	UserStatusEnable   uint = 1
)

type User struct {
	gorm.M
	Mobile     string          `json:"mobile" gorm:"index:idx_mobile,unique;size:20;comment:手机号码"`
	Password   string          `json:"password" gorm:"size:64;comment:密码"`
	Avatar     string          `json:"avatar" gorm:"comment:头像地址"`
	Name       string          `json:"name" gorm:"size:64;comment:姓名"`
	Status     *uint           `json:"status" gorm:"type:tinyint(1);default:1;comment:账号状态(0-禁用 1-启用)"`
	RoleId     uint            `json:"roleId" gorm:"comment:角色ID"`
	Role       Role            `json:"role" gorm:"foreignKey:RoleId"`
	LastLogin  carbon.DateTime `json:"lastLogin" gorm:"comment:最后登录时间"`
	Locked     uint            `json:"locked" gorm:"type:tinyint(1);default:0;comment:账户锁定(0-未锁定 1-已锁定)"`
	LockExpire int64           `json:"lockExpire" gorm:"comment:账户锁定剩余时间"`
	PwdWrong   int             `json:"pwdWrong" gorm:"comment:密码错误次数"`
}
