package gorm

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Conf struct {
	Type        string `json:"Type"`        // type of database such as mysql or postgres
	Host        string `json:"Host"`        // database address
	Port        int    `json:"Port"`        // database port
	Config      string `json:"Config"`      // extra config such as charset=utf8mb4&parseTime=True&loc=Local
	DbName      string `json:"DbName"`      // database name
	Username    string `json:"Username"`    // database username
	Password    string `json:"Password"`    // database password
	MaxIdleConn int    `json:"MaxIdleConn"` // the maximum number of connections in the idle connection pool
	MaxOpenConn int    `json:"MaxOpenConn"` // the maximum number of open connections to the database
	LogMode     string `json:"LogMode"`     // set gorm global log mode
}

func (g *Conf) MysqlDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", g.Username, g.Password, g.Host, g.Port, g.DbName, g.Config)
}

func (g *Conf) PostgresDsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d %s", g.Host, g.Username, g.Password, g.DbName, g.Port, g.Config)
}

func (g *Conf) NewGormDb() (*gorm.DB, error) {
	switch g.Type {
	case "mysql":
		return Mysql(g)
	case "postgres":
		return Postgres(g)
	default:
		return Mysql(g)
	}
}

func Mysql(g *Conf) (*gorm.DB, error) {
	if g.DbName == "" {
		return nil, errors.New("database name cannot be empty")
	}
	mysqlConfig := mysql.Config{
		DSN:                       g.MysqlDsn(),
		DefaultStringSize:         256,   // default size of string fields
		DisableDatetimePrecision:  true,  // disable datetime precision which not supported before mysql 5.6
		DontSupportRenameIndex:    true,  // drop and create when rename index which not supported before mysql 5.7 and mariadb
		SkipInitializeWithVersion: false, // autoconfiguration based on current mysql version
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: logger.New(Logger{}, logger.Config{
			SlowThreshold:             1 * time.Second, // sql slow query threshold
			Colorful:                  false,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  getLevel(g.LogMode),
		}),
	}); err != nil {
		return nil, err
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(g.MaxIdleConn)
		sqlDB.SetMaxOpenConns(g.MaxOpenConn)
		return db, nil
	}
}

func Postgres(g *Conf) (*gorm.DB, error) {
	if g.DbName == "" {
		return nil, errors.New("database name cannot be empty")
	}
	postgresConfig := postgres.Config{
		DSN:                  g.PostgresDsn(),
		PreferSimpleProtocol: false, // disable implicit prepared statement usage
	}

	if db, err := gorm.Open(postgres.New(postgresConfig), &gorm.Config{
		Logger: logger.New(Logger{}, logger.Config{
			SlowThreshold:             1 * time.Second,
			Colorful:                  false,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  getLevel(g.LogMode)}),
	}); err != nil {
		return nil, err
	} else {
		sqlDb, _ := db.DB()
		sqlDb.SetMaxIdleConns(g.MaxIdleConn)
		sqlDb.SetMaxOpenConns(g.MaxOpenConn)
		return db, nil
	}
}

func getLevel(logMode string) (l logger.LogLevel) {
	switch logMode {
	case "info":
		l = logger.Info
	case "warn":
		l = logger.Warn
	case "error":
		l = logger.Error
	case "silent":
		l = logger.Silent
	default:
		l = logger.Error
	}
	return l
}
