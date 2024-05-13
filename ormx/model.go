package ormx

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type OrmType string

const (
	MySQLType OrmType = "mysql"
)

type OrmxConfig struct {
	Type        OrmType
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
	MaxIdleCons int
	MaxOpenCons int
	LogMode     string
}

func Setup(config *OrmxConfig) (db *gorm.DB, err error) {
	var ormLogger logger.Interface

	if config.LogMode == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default.LogMode(logger.Error)
	}

	switch config.Type {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User,
			config.Password, config.Host, config.Name)

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   config.TablePrefix,
				SingularTable: true,
			},
			Logger: ormLogger,
		})
		if err != nil {
			return nil, fmt.Errorf("gorm open db conncection fail, err: %w", err)
		}
	default:
		return nil, fmt.Errorf("unknown database type: %s", config.Type)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get db fail, err: %w", err)
	}

	sqlDB.SetMaxIdleConns(config.MaxIdleCons)
	sqlDB.SetMaxOpenConns(config.MaxOpenCons)

	return
}
