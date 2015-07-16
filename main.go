package main

import (
	"github.com/yanjinzh6/youzhe-server/db"
	"github.com/yanjinzh6/youzhe-server/server"
)

func main() {
	server.InitServer()
	db.InitRedis()
	defer db.Close()
}
