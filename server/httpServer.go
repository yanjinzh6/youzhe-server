package server

import (
	"reflect"
	"net/http"
	"html/template"
	"github.com/yanjinzh6/youzhe-server/conf"
	"github.com/yanjinzh6/youzhe-server/tools"
)

func InitServer() {
	http.Handle("/css/", http.FileServer(http.Dir("template")))
	http.Handle("/js/", http.FileServer(http.Dir("template")))
	http.Handle("/img/", http.FileServer(http.Dir("template")))
	http.Handle("/fonts/", http.FileServer(http.Dir("template")))
	config := conf.InitConfig(tools.DEFAULT_CONFIG_FILE)
	in := make([]reflect.Value, 2)
	//http.HandleFunc(config.Server.HandlerList[0].Action, config.Server.HandlerList[0].MyFunc)
	for _, ht := range config.Server.HandlerList {
		med, err := conf.FindMethod(ht.MyFunc)
		if err == nil {
			handler := func (w http.ResponseWriter, r *http.Request) {
				in[0] = reflect.ValueOf(w)
				in[1] = reflect.ValueOf(r)
				//tools.Println(in, "\n", w, "\n", r)
				med.Call(in)
			}
			http.HandleFunc(ht.Action, handler)
		} else {
			tools.Println(err)
		}
	}
	http.HandleFunc("/", NotFoundHandler)
	err := http.ListenAndServe(":" + config.Server.Port, nil)
	if (err != nil) {
		tools.Println(err)
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		t, err := template.ParseFiles("template/index.html")
		if err != nil {
			tools.Println(err)
		} else {
			t.Execute(w, nil)
		}
	} else {
		t, err := template.ParseFiles("template/404.html")
		if err != nil {
			tools.Println(err)
		} else {
			t.Execute(w, nil)
		}
	}
}