package db

import (
	"api/config"
	"api/logs"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"time"
)

var DB *gorm.DB

func init() {
	DB = NewMysqlGorm(config.Evn.Database)
}

func NewMysqlGorm(url string) *gorm.DB {
	var err error
	var DB *gorm.DB
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: url,
	}), Config())
	sqlDB, err := DB.DB()
	if err != nil {
		zap.L().Error(err.Error())
		return nil
	}
	if err = sqlDB.Ping(); err != nil {
		zap.L().Error(err.Error())
		return nil
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return DB
}

func NewPostgresqlGorm(url string) *gorm.DB {
	var err error
	var DB *gorm.DB
	DB, err = gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  url,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), Config())
	sqlDB, err := DB.DB()
	if err != nil {
		zap.L().Error(err.Error())
		return nil
	}
	if err = sqlDB.Ping(); err != nil {
		zap.L().Error(err.Error())
		return nil
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return DB
}

func Config() *gorm.Config {
	newLogger := logger.New(
		NewGLog(),
		logger.Config{
			SlowThreshold:             time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,      // Log level
			IgnoreRecordNotFoundError: false,            // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,            // Don't include params in the SQL log
			Colorful:                  false,            // Disable color
		},
	)
	return &gorm.Config{
		Logger:                                   newLogger,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   true, // skip the snake_casing of names
		},
	}
}

type GLog struct {
	*zap.Logger
}

func (l GLog) Printf(format string, args ...interface{}) {
	sprintf := "\r" + fmt.Sprintf(format, args...)
	l.Info(sprintf)
}

func NewGLog() *GLog {
	return &GLog{logs.Log.WithOptions(zap.AddCallerSkip(4))}
}
