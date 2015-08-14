package dbm

import (
	"github.com/yanjinzh6/youzhe-server/tools"
	"testing"
)

type User struct {
	Id    string
	Name  string
	Email string
}

func TestInitMongodb(t *testing.T) {
	InitMongodb()
	q, err := MyMongodb.Db.C("users").RemoveAll(nil)
	tools.Println(q, err, q.Removed)
	// u := User{
	// 	Id:    "12345",
	// 	Name:  "User1",
	// 	Email: "User1@yozh.com",
	// }
	// err := MyMongodb.Db.C("users").Insert(&u)
	tools.Println(err)
	n, err := MyMongodb.Db.C("users").Count()
	tools.Println(n, err)
	var us []User
	err = MyMongodb.Db.C("users").Find(nil).Skip(1).Limit(1).All(&us)
	tools.Println(us, err)
}
