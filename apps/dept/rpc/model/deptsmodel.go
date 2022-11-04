package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DeptsModel = (*customDeptsModel)(nil)

type (
	// DeptsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeptsModel.
	DeptsModel interface {
		deptsModel
	}

	customDeptsModel struct {
		*defaultDeptsModel
	}
)

// NewDeptsModel returns a model for the database table.
func NewDeptsModel(conn sqlx.SqlConn, c cache.CacheConf) DeptsModel {
	return &customDeptsModel{
		defaultDeptsModel: newDeptsModel(conn, c),
	}
}
