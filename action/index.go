package action

import (
	"fmt"
	"github.com/yanjinzh6/youzhe-server/tools"
	"html/template"
	"net/http"
)

type Obje struct {
	name string
}

func init() {
	i := &Obje{
		name: "index",
	}
	MyAcMap.Set(i.name, i)
}

func (i *Obje) Index(w http.ResponseWriter, r *http.Request) {
	//http.Redirect(w, r, "template/index.html", http.StatusFound)
	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		tools.Println(err)
	} else {
		t.Execute(w, nil)
	}
}

func (i *Obje) Login(w http.ResponseWriter, r *http.Request) {
	//http.Redirect(w, r, "template/index.html", http.StatusFound)
	t, err := template.ParseFiles("template/login.html")
	if err != nil {
		tools.Println(err)
	} else {
		t.Execute(w, nil)
	}
}

func (i *Obje) Demo(w http.ResponseWriter, r *http.Request) {
	tools.Println(w.Header(), r.Host)
	params := r.URL.Query()
	id := params.Get(":id")
	fmt.Fprintf(w, "{'id': %s}", id)
}
