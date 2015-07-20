package action

import (
	"github.com/yanjinzh6/youzhe-server/dbm"
	"github.com/yanjinzh6/youzhe-server/tools"
	"gopkg.in/mgo.v2/bson"
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
	t := r.FormValue("collectiontype")
	dbm.MyMongodb.Db.C(t).Insert(r.Form)
	n, err := dbm.MyMongodb.Db.C(t).Count()
	tools.Println("user:", n, err)
	q := dbm.MyMongodb.Db.C(t).Find(bson.M{"collectiontype": "user"})
	tools.Println("user:", q)
	http.Redirect(w, r, "addUserPage", http.StatusFound)
}
