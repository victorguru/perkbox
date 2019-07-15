package main

import (
	"database/sql"
	"log"

	"github.com/DavidHuie/gomigrate"
)

// Database holds the database connection
type Database int

// ConnectDB will create a db object
func (db *Database) ConnectDB() *sql.DB {
	log.Println("----connectDB")
	conf := getConf()

	conn, err := sql.Open("mysql", conf.DSNDB)
	if err != nil {
		log.Println("----connectDB error1")
		log.Println(err)
		panic(conn)
	}
	// defer conn.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = conn.Ping()
	if err != nil {
		log.Println("----connectDB error2")
		log.Println(err)
		panic(conn)
	}

	// Create table and start with some data
	migrator, err := gomigrate.NewMigrator(conn, gomigrate.Mysql{}, "./migrations")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	err = migrator.Migrate()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return conn
}
