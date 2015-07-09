package action

import (
	"net/http"
	"github.com/yanjinzh6/youzhe-server/tools"
)

type index struct {
	name string
}

func init() {
	i := &index{
		name : "index",
	}
	MyAcMap.Set(i.name, i)
}

func (i *index) Index(w http.ResponseWriter, r *http.Request) {
	tools.Println(w.Header(), r.Host)
}

func (i *index) Demo(w http.ResponseWriter, r *http.Request) {
	tools.Println(w.Header(), r.Host)
}