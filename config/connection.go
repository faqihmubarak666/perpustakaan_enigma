package config

import (
	"database/sql"
	"fmt"
	"gomux/utils"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {

	dbUser := utils.ReadEnv("dbUser", "root")
	dbPass := utils.ReadEnv("dbPass", "")
	dbHost := utils.ReadEnv("dbHost", "localhost")
	dbPort := utils.ReadEnv("dbPort", "3306")
	dbName := utils.ReadEnv("dbName", "")

	loadData := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, _ := sql.Open("mysql", loadData)

	err := db.Ping()
	if err != nil {
		log.Print(err)
	}
	return db, nil
}
