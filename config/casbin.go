package config

import (
	"errors"
	"fiber-g/pkg/gorm"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/zeromicro/go-zero/core/logx"
)

type CasbinConf struct {
	ModelConf string
}

func (c CasbinConf) NewCasbin(g gorm.Conf) (*casbin.Enforcer, error) {
	var dsn string
	switch g.Type {
	case "mysql":
		dsn = g.MysqlDsn()
	case "postgres":
		dsn = g.PostgresDsn()
	default:
		return nil, errors.New("unsupported database type")
	}

	adapter, err := gormadapter.NewAdapter(g.Type, dsn)
	logx.Must(err)

	var text string
	if c.ModelConf == "" {
		text = `
				[request_definition]
				r = sub, obj, act
				
				[policy_definition]
				p = sub, obj, act
				
				[role_definition]
				g = _, _
				
				[policy_effect]
				e = some(where (p.eft == allow))
				
				[matchers]
				m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
               `
	} else {
		text = c.ModelConf
	}

	m, err := model.NewModelFromString(text)
	logx.Must(err)

	enforcer, err := casbin.NewEnforcer(m, adapter)
	logx.Must(err)

	err = enforcer.LoadPolicy()
	logx.Must(err)

	return enforcer, nil
}
