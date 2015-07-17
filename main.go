package main

import (
	"github.com/yanjinzh6/youzhe-server/db"
	"github.com/yanjinzh6/youzhe-server/dbm"
	"github.com/yanjinzh6/youzhe-server/server"
)

func main() {
	dbm.InitMongodb()
	db.InitRedis()
	server.InitServer()

	defer db.Close()
	defer dbm.Close()
}
