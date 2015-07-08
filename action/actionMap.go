package action

import (
	"github.com/yanjinzh6/youzhe-server/tools"
)

type actionMap struct {
	acMap map[string]interface{}
}

var myAcMap *actionMap

func init() {
	myAcMap = New()
}

func New() (aMap *actionMap) {
	if (myAcMap == nil) {
		myAcMap = &actionMap {
			acMap : make(map[string]interface{}),
		}
	}
	return myAcMap
}

func (a *actionMap) get(key string) (val interface{}) {
	return a.acMap[key]
}

func (a *actionMap) set(key string, val interface{}) (oldVal interface{}, err error) {
	if a.acMap[key] != nil {
		oldVal = a.acMap[key]
		err = tools.KeyNotNilError
	}
	a.acMap[key] = val
	return
}