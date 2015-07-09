package server

import (
	"reflect"
	"net/http"
	"github.com/yanjinzh6/youzhe-server/conf"
	"github.com/yanjinzh6/youzhe-server/tools"
)

func InitServer() {
	config := conf.InitConfig(tools.DEFAULT_CONFIG_FILE)
	in := make([]reflect.Value, 2)
	//http.HandleFunc(config.Server.HandlerList[0].Action, config.Server.HandlerList[0].MyFunc)
	for _, ht := range config.Server.HandlerList {
		med, err := conf.FindMethod(ht.MyFunc)
		if err == nil {
			handler := func (w http.ResponseWriter, r *http.Request) {
				in[0] = reflect.ValueOf(w)
				in[1] = reflect.ValueOf(r)
				tools.Println(in, "\n", w, "\n", r)
				med.Call(in)
			}
			http.HandleFunc(ht.Action, handler)
		} else {
			tools.Println(err)
		}
	}
	err := http.ListenAndServe(":" + config.Server.Port, nil)
	if (err != nil) {
		tools.Println(err)
	}
}