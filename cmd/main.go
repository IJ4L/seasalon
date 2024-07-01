package main

import (
	"database/sql"
	server "gitlab/go-prolog-api/example/internal"
	"gitlab/go-prolog-api/example/repository"
	"gitlab/go-prolog-api/example/util"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	user := repository.NewProduk(conn)
	server, err := server.NewServer(config, user)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}