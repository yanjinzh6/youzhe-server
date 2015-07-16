package server

import (
	"github.com/yanjinzh6/youzhe-server/conf"
	"github.com/yanjinzh6/youzhe-server/log"
	"github.com/yanjinzh6/youzhe-server/tools"
	"html/template"
	"net/http"
	"reflect"
	"strings"
)

func InitServer() {
	http.Handle("/css/", http.FileServer(http.Dir("template")))
	http.Handle("/js/", http.FileServer(http.Dir("template")))
	http.Handle("/img/", http.FileServer(http.Dir("template")))
	http.Handle("/fonts/", http.FileServer(http.Dir("template")))
	//config := conf.InitConfig(tools.DEFAULT_CONFIG_FILE)
	in := make([]reflect.Value, 2)
	//http.HandleFunc(config.Server.HandlerList[0].Action, config.Server.HandlerList[0].MyFunc)
	for _, ht := range conf.MyConfig.Server.HandlerList {
		med, err := conf.FindMethod(ht.MyFunc)
		if err == nil {
			handler := func(w http.ResponseWriter, r *http.Request) {
				tools.Println("RemoteAddr:", r.RemoteAddr, "ServerHost:", r.Host, "RequestURI", r.RequestURI)
				var nUrl = ""
				if strings.Contains(r.Host, ":") {
					var hp = strings.Split(r.Host, ":")
					if strings.EqualFold(hp[1], conf.MyConfig.Server.Port) {
						nUrl = "https://" + hp[0] + ":" + conf.MyConfig.Server.Ports + r.RequestURI
					}
				} else {
					nUrl = "https://" + r.Host + ":" + conf.MyConfig.Server.Ports + r.RequestURI
				}
				if nUrl == "" {
					in[0] = reflect.ValueOf(w)
					in[1] = reflect.ValueOf(r)
					//tools.Println(in, "\n", w, "\n", r)
					med.Call(in)
				} else {
					http.Redirect(w, r, nUrl, http.StatusFound)
				}
			}
			http.HandleFunc(ht.Action, handler)
		} else {
			tools.Println(err)
		}
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tools.Println("RemoteAddr:", r.RemoteAddr, "ServerHost:", r.Host, "RequestURI", r.RequestURI)
		var nUrl = ""
		if strings.Contains(r.Host, ":") {
			var hp = strings.Split(r.Host, ":")
			if strings.EqualFold(hp[1], conf.MyConfig.Server.Port) {
				nUrl = "https://" + hp[0] + ":" + conf.MyConfig.Server.Ports + r.RequestURI
			}
		} else {
			nUrl = "https://" + r.Host + ":" + conf.MyConfig.Server.Ports + r.RequestURI
		}
		if nUrl == "" {
			if r.URL.Path == "/" {
				t, err := template.ParseFiles("template/index.html")
				if err != nil {
					tools.Println(err)
				} else {
					t.Execute(w, nil)
				}
			} else if r.URL.Path == "/favicon.ico" {
				http.ServeFile(w, r, "template/favicon.ico")
			} else {
				t, err := template.ParseFiles("template/404.html")
				if err != nil {
					tools.Println(err)
				} else {
					t.Execute(w, nil)
				}
			}
		} else {
			http.Redirect(w, r, nUrl, http.StatusFound)
		}
	})
	go func() {
		err := http.ListenAndServeTLS(":"+conf.MyConfig.Server.Ports, "cert.pem", "key.pem", nil)
		if err != nil {
			log.Fatalln("ListenAndServeTLS error: ", err)
		}
	}()
	err := http.ListenAndServe(":"+conf.MyConfig.Server.Port, nil)
	if err != nil {
		log.Fatalln("ListenAndServe error: ", err)
	}
}
