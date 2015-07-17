package main

import (
	"github.com/yanjinzh6/youzhe-server/db"
	"github.com/yanjinzh6/youzhe-server/server"
)

func main() {
	db.InitRedis()
	server.InitServer()

	defer db.Close()
}
