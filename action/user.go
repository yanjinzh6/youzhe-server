package action

import (
	// "github.com/yanjinzh6/youzhe-server/dbm"
	"github.com/yanjinzh6/youzhe-server/tools"
	"html/template"
	"net/http"
)

func init() {
	i := &Obje{
		name: "user",
	}
	MyAcMap.Set(i.name, i)
}

func (i *Obje) AddUserPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/user/user.html")
	if err != nil {
		tools.Println(err)
	} else {
		t.Execute(w, nil)
	}
}

func (i *Obje) AddUser(w http.ResponseWriter, r *http.Request) {
	tools.Println(r.FormValue("collectiontype"), r.ParseForm())
	i.AddUserPage(w, r)
}
