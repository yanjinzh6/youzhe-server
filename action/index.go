package action

import (
	"net/http"
	"html/template"
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
	//http.Redirect(w, r, "template/index.html", http.StatusFound)
	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		tools.Println(err)
	} else {
		t.Execute(w, nil)
	}
}

func (i *index) Demo(w http.ResponseWriter, r *http.Request) {
	tools.Println(w.Header(), r.Host)
}