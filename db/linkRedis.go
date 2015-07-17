package db

import (
	"github.com/garyburd/redigo/redis"
	"github.com/yanjinzh6/youzhe-server/conf"
	"github.com/yanjinzh6/youzhe-server/log"
	"github.com/yanjinzh6/youzhe-server/tools"
)

type redisConn struct {
	Conn redis.Conn
}

var MyRedis *redisConn

func InitRedis() *redisConn {
	tools.Println(&MyRedis, MyRedis)
	if MyRedis == nil {
		rconn, err := redis.Dial("tcp", conf.MyConfig.Redis.Addr+":"+conf.MyConfig.Redis.Port)
		if err != nil {
			log.Println("link redis error", err)
		}
		MyRedis = &redisConn{
			Conn: rconn,
		}
		log.Println("redis connect success!")
	}
	return MyRedis
}

func (r *redisConn) close() {
	r.Conn.Close()
}

func Close() {
	tools.Println(&MyRedis, MyRedis)
	if MyRedis != nil && MyRedis.Conn != nil {
		MyRedis.close()
		log.Println("close redis connect!")
	} else {
		log.Println("didn't found redis connect!")
	}
}
