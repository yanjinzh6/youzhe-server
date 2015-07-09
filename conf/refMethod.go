package conf

import (
	"reflect"
	"strings"
	"github.com/yanjinzh6/youzhe-server/action"
	"github.com/yanjinzh6/youzhe-server/tools"
)

func FindMethod(f string) (m reflect.Value, err error) {
	var temp = strings.Split(f, "/")
	var aStr [2]string
	flag := 0
	for _, s := range temp {
		if s != "" {
			aStr[flag] = s
			flag ++
		}
	}
	tools.Println("actionStr is: ", aStr, len(aStr), aStr[0], aStr[1])
	if (len(aStr) == 2 && aStr[0] != "" && aStr[1] != "") {
		var ac = action.MyAcMap.Get(aStr[0])
		tools.Println("action struct is: ", ac)
		if ac != nil {
			t := reflect.ValueOf(ac)
			m = t.MethodByName(aStr[1])
			tools.Println("action method is: ", m, m.IsValid())
			if !m.IsValid() {
				err = tools.FuncNotFoundError
			}
		} else {
			err = tools.ActionNotFoundError
		}
	} else {
		err = tools.ActionParamError
	}
	return
}