package server

import (
	"net/http"
	"github.com/yanjinzh6/youzhe-server/conf"
	"github.com/yanjinzh6/youzhe-server/tools"
)

func InitServer() {
	config := conf.InitConfig(tools.DEFAULT_CONFIG_FILE)
	//http.HandleFunc(config.Server.HandlerList[0].Action, config.Server.HandlerList[0].MyFunc)
	http.HandleFunc(config.Server.HandlerList[0].Action, index)
	err := http.ListenAndServe(":" + config.Server.Port, nil)
	if (err != nil) {
		tools.Println(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	tools.Println(w, r)
}