package session

import (
	"github.com/yanjinzh6/youzhe-server/tools"
	"testing"
)

func TestGet(t *testing.T) {
	s := SessionManage{ObjId: "123", UserName: "user"}
	tools.Println(Set(s))
	tools.Println(Get("123"))
}
