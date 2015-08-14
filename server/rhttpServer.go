package server

import (
	"github.com/drone/routes"
	"github.com/yanjinzh6/youzhe-server/action"
	"github.com/yanjinzh6/youzhe-server/conf"
	"github.com/yanjinzh6/youzhe-server/log"
	"github.com/yanjinzh6/youzhe-server/tools"
	"net/http"
	// "os"
	"reflect"
	"strings"
)

func Create() {
	http.Handle("/", muxConfig())
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

func muxConfig() *routes.RouteMux {
	mux := routes.New()
	mux.Static("/css/", "template")
	mux.Static("/js/", "template")
	mux.Static("/img/", "template")
	mux.Static("/fonts/", "template")
	mux.Static("/views/", "template")
	mux.Static("/less/", "template")
	mux.Static("/favicon.ico", "template")
	mux.Static("/apple-touch-icon.png", "template")
	mux.Static("/tile.png", "template")
	mux.Static("/tile-wide.png", "template")
	mux.Static("/browserconfig.xml", "template")
	mux.Static("/crossdomain.xml", "template")
	mux.Static("/humans.txt", "template")
	mux.Static("/LICENSE.txt", "template")
	mux.Static("/robots.txt", "template")
	// mux.Static("/", "template")
	in := make([]reflect.Value, 2)
	//http.HandleFunc(config.Server.HandlerList[0].Action, config.Server.HandlerList[0].MyFunc)
	for _, ht := range conf.MyConfig.Server.HandlerList {
		med, err := action.FindMethod(ht.MyFunc)
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
			mux.AddRoute(ht.Method, ht.Action, handler)
			tools.Println("addroute:", ht.Method, ht.Action, med)
		} else {
			tools.Println(err)
		}
	}
	return mux
}
