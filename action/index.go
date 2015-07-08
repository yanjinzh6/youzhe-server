package action

import (
	"net/http"
	"github.com/yanjinzh6/youzhe-server/tools"
)

type index struct {}

func (i *index) index(w http.ResponseWriter, r *http.Request) {
	tools.Println(w, r)
}