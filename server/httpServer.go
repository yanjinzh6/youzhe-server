package server

import (
	"net/http"
)

type MyServer struct {
	Addr string
	Port string
	HandlerList handlerItem[]
}

type handlerItem struct {
	Path string
	function string
}