package main

import (
	"sosmed/config"
	"sosmed/utils/database"
)

func main() {
	var dbConfig = new(config.Database)
	if err := dbConfig.LoadFromEnv(); err != nil {
		panic(err)
	}

	dbConnection, err := database.MysqlInit(*dbConfig)
	if err != nil {
		panic(err)
	}

	if err := database.MysqlMigrate(dbConnection); err != nil {
		panic(err)
	}
}
