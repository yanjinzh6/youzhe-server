package session

import (
	"bytes"
	"encoding/json"
	"github.com/yanjinzh6/youzhe-server/db"
	"github.com/yanjinzh6/youzhe-server/tools"
)

type sessionManage struct {
	objId    string
	userName string
}

func Get(id string) (s sessionManage, e error) {
	buf := make([]byte, 0, tools.DEFAULT_BUFFER_SIZE)
	bufs := bytes.NewBuffer(buf)
	val, err := db.MyRedisPool.Pool.Get.Do("GET", id)
	tools.Println(val, err)
	n, err := bufs.WriteTo(val)
	tools.Println(n, err)
	err = json.Unmarshal(bufs.Bytes(), &sessionManage{})
	tools.Println(err)
}
