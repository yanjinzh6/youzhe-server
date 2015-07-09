package action

import (
	"github.com/yanjinzh6/youzhe-server/tools"
)

type ActionMap struct {
	acMap map[string]interface{}
}

var MyAcMap *ActionMap

func init() {
	MyAcMap = New()
}

func New() (aMap *ActionMap) {
	if (MyAcMap == nil) {
		MyAcMap = &ActionMap {
			acMap : make(map[string]interface{}),
		}
	}
	return MyAcMap
}

func (a *ActionMap) Get(key string) (val interface{}) {
	return a.acMap[key]
}

func (a *ActionMap) Set(key string, val interface{}) (oldVal interface{}, err error) {
	if a.acMap[key] != nil {
		oldVal = a.acMap[key]
		err = tools.KeyNotNilError
	}
	a.acMap[key] = val
	tools.Println("ActionMap size: ", len(a.acMap))
	return
}