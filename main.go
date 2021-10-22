package main

import (
	"strconv"
	"voting-system/pkg/configuration"
	"voting-system/pkg/database"
	"voting-system/pkg/exception"
)

func main() {
	config := configuration.NewConfigurationImpl(".env")

	poolMinConfig, err := strconv.Atoi(config.Get("MYSQL_POOL_MIN"))
	exception.PanicIfError(err)
	poolMaxConfig, err := strconv.Atoi(config.Get("MYSQL_POOL_MAX"))
	exception.PanicIfError(err)
	dbConfig := configuration.DbConfig{
		User:     config.Get("MYSQL_USER"),
		Password: config.Get("MYSQL_PASSWORD"),
		Host:     config.Get("MYSQL_HOST"),
		Port:     config.Get("MYSQL_PORT"),
		Database: config.Get("MYSQL_DATABASE"),
		PoolMin:  poolMinConfig,
		PoolMax:  poolMaxConfig,
	}

	_ = database.NewDatabase(dbConfig)
}
