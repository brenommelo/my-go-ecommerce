package main

import (
	"database/sql"
	"log"

	"github.com/brenommelo/my-go-ecommerce/cmd/api"
	configs "github.com/brenommelo/my-go-ecommerce/config"
	"github.com/brenommelo/my-go-ecommerce/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewMysqlStorage(mysql.Config{
		User:                 configs.Envs.DBUser,
		Passwd:               configs.Envs.DBPassword,
		Addr:                 configs.Envs.DBAddress,
		DBName:               configs.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Sucessfully connected!")
}
