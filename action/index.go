package action

import (
	"github.com/yanjinzh6/youzhe-server/db"
	"github.com/yanjinzh6/youzhe-server/tools"
	"html/template"
	"net/http"
)

type index struct {
	name string
}

func init() {
	i := &index{
		name: "index",
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

func (i *index) Login(w http.ResponseWriter, r *http.Request) {
	//http.Redirect(w, r, "template/index.html", http.StatusFound)
	n, err := db.MyRedis.Conn.Do("exists", "session")
	tools.Println("login", n, err)
	t, err := template.ParseFiles("template/login.html")
	if err != nil {
		tools.Println(err)
	} else {
		t.Execute(w, nil)
	}
}

func (i *index) Demo(w http.ResponseWriter, r *http.Request) {
	tools.Println(w.Header(), r.Host)
}
