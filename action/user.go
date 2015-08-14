package action

import (
	"encoding/json"
	"fmt"
	"github.com/yanjinzh6/youzhe-server/dbm"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

type User struct {
	Id    bson.ObjectId `bson: "_id"`
	Name  string        `bson: "name"`
	Email string        `bson: "email"`
}

func init() {
	i := &Obje{
		name:           "user",
		collectionName: "users",
	}
	MyAcMap.Set(i.name, i)
}

func (i *Obje) UserList(rw http.ResponseWriter, req *http.Request) {
	var users []User
	// u := User{
	// 	Id:    bson.NewObjectId(),
	// 	Name:  "User1",
	// 	Email: "User1@yozh.com",
	// }
	// u2 := User{
	// 	Id:    bson.NewObjectId(),
	// 	Name:  "User2",
	// 	Email: "User2@yozh.com",
	// }
	// users = append(users, u)
	// users = append(users, u2)
	params := req.URL.Query()
	pageSize, _ := strconv.Atoi(params.Get("pagesize"))
	pageIndex, _ := strconv.Atoi(params.Get("pageindex"))
	fmt.Println(pageIndex, pageSize)
	err := dbm.MyMongodb.Db.C(i.collectionName).Find(nil).Skip(pageSize * pageIndex).Limit(pageSize).All(&users)
	// err := dbm.MyMongodb.Db.C(i.collectionName).Find(nil).All(&users)
	fmt.Println(users, err)
	var rs RespState
	if len(users) != 0 {
		rs.State = "success"
		rs.Content = users

	} else {
		rs.State = "none"
		rs.Content = "can not find user"
	}
	jusers, err := json.Marshal(rs)
	fmt.Println(rs, jusers, err)
	// fmt.Fprint(rw, jusers)
	rw.Write(jusers)
}

func (i *Obje) QueryUser(rw http.ResponseWriter, req *http.Request) {
	// u := User{
	// 	Id:    bson.NewObjectId(),
	// 	Name:  "User1",
	// 	Email: "User1@yozh.com",
	// }
	var user User
	params := req.URL.Query()
	id := params.Get(":id")
	err := dbm.MyMongodb.Db.C(i.collectionName).FindId(id).One(&user)
	fmt.Println(err)
	var juser []byte
	if user.Id != "" {
		juser, _ = json.Marshal(user)
		// fmt.Println(err)
	} else {
		// juser = "{\"state\": \"none\"}"
	}
	// fmt.Fprint(rw, juser)
	rw.Write(juser)
}

func (i *Obje) AddUser(rw http.ResponseWriter, req *http.Request) {
	user := User{
		Id:    bson.NewObjectId(),
		Name:  req.FormValue("name"),
		Email: req.FormValue("email"),
	}
	err := dbm.MyMongodb.Db.C(i.collectionName).Insert(&user)
	juser, err := json.Marshal(user)
	fmt.Println(err)
	// fmt.Fprint(rw, juser)
	rw.Write(juser)
}

func (i *Obje) UpdateUser(rw http.ResponseWriter, req *http.Request) {
	u := User{
		Id:    bson.NewObjectId(),
		Name:  "User1",
		Email: "User1@yozh.com",
	}
	juser, err := json.Marshal(u)
	fmt.Println(err)
	// fmt.Fprint(rw, juser)
	rw.Write(juser)
}

func (i *Obje) DeleteUser(rw http.ResponseWriter, req *http.Request) {
	u := User{
		Id:    bson.NewObjectId(),
		Name:  "User1",
		Email: "User1@yozh.com",
	}
	juser, err := json.Marshal(u)
	fmt.Println(err)
	// fmt.Fprint(rw, juser)
	rw.Write(juser)
}
