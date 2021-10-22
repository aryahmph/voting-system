package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
	"voting-system/pkg/configuration"
	"voting-system/pkg/exception"
)

func NewDatabase(config configuration.DbConfig) *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	db, err := sqlx.Connect("mysql", dsn)
	exception.PanicIfError(err)

	db.SetMaxIdleConns(config.PoolMin)
	db.SetMaxOpenConns(config.PoolMax)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
