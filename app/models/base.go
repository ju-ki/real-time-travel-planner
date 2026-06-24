package models

import (
	"database/sql"
	"log"
	"real-time-travel-planner/config"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

func init() {
	var err error
	DbConnection, err = sql.Open(config.AppConfig.SQLDriver, config.AppConfig.DbName)
	if err != nil {
		log.Fatalln(err)
		return
	}

	_, err = DbConnection.Exec(`
	    CREATE TABLE IF NOT EXISTS USERS  (
			userId INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatalln(err)
		return
	}

	_, err = DbConnection.Exec(`
	    CREATE TABLE IF NOT EXISTS ROOMS (
			roomId INTEGER PRIMARY KEY AUTOINCREMENT,
			roomName TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Println(err)
		return
	}

	rows, err := DbConnection.Query("SELECT name FROM sqlite_master WHERE type='table'")
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	var tableName string
	log.Println("Tables in SQLite:")

	// 結果をループ処理して出力
	for rows.Next() {
		err := rows.Scan(&tableName)
		if err != nil {
			log.Println(err)
		}
		log.Println("-", tableName)
	}

	log.Println("Database connection established and tables created successfully.")
}
