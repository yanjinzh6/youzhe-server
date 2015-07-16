package db

import (
	"github.com/garyburd/redigo/redis"
	"github.com/yanjinzh6/youzhe-server/conf"
	"github.com/yanjinzh6/youzhe-server/log"
)

type redisConn struct {
	Conn redis.Conn
}

var MyRedis *redisConn

func InitRedis() *redisConn {
	if MyRedis.Conn == nil {
		rconn, err := redis.Dial("tcp", ":"+conf.MyConfig.Redis.Port)
		if err != nil {
			log.Println("link redis error", err)
		}
		MyRedis = &redisConn{
			Conn: rconn,
		}
	}
	return MyRedis
}

func (r *redisConn) close() {
	r.Conn.Close()
}

func Close() {
	MyRedis.close()
	log.Println("close redis connect!")
}
