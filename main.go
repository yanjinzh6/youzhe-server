package main

import (
	// "bufio"
	"github.com/yanjinzh6/youzhe-server/db"
	"github.com/yanjinzh6/youzhe-server/dbm"
	"github.com/yanjinzh6/youzhe-server/server"
	_ "github.com/yanjinzh6/youzhe-server/signal"
	// "github.com/yanjinzh6/youzhe-server/tools"
	// "os"
)

func main() {
	dbm.InitMongodb()
	db.InitRedis()
	// server.InitServer()
	server.Create()

	defer db.Close()
	defer dbm.Close()

	// stop := false
	// reader := bufio.NewReader(os.Stdin)
	// for !stop {
	// 	data, _, _ := reader.ReadLine()
	// 	command := string(data)
	// 	if command == "stop" {
	// 		stop = true
	// 	}
	// 	tools.Println("command:", command)
	// }
}
