package session

import (
	"github.com/garyburd/redigo/redis"
	"github.com/yanjinzh6/youzhe-server/db"
	"github.com/yanjinzh6/youzhe-server/tools"
)

type SessionManage struct {
	ObjId    string `redis:"objId"`
	UserName string `redis:"userName"`
}

func Get(id string) (s SessionManage, e error) {
	conn := db.MyRedisPool.Pool.Get()
	defer conn.Close()
	val, err := conn.Do("HGETALL", id)
	tools.Println("session manage", val, err)
	vals, err := db.Values(val, err)
	err = redis.ScanStruct(vals, &s)
	tools.Println("session manage", err)
	return
}

func Set(s SessionManage) (err error) {
	conn := db.MyRedisPool.Pool.Get()
	defer conn.Close()
	_, err = conn.Do("HMSET", redis.Args{}.Add(s.ObjId).AddFlat(&s)...)
	tools.Println("session manage", err)
	return
}
