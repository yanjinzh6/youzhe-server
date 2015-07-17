package db

import (
	"github.com/garyburd/redigo/redis"
	"github.com/yanjinzh6/youzhe-server/conf"
	"github.com/yanjinzh6/youzhe-server/log"
	"github.com/yanjinzh6/youzhe-server/tools"
	"time"
)

type redisConn struct {
	Conn redis.Conn
}

var MyRedis *redisConn

type redisPool struct {
	Pool *redis.Pool
}

var MyRedisPool *redisPool

// func init() {
// 	InitRedis()
// }

func NewRedis(server, password string) (myredis *redisConn, e error) {
	if MyRedis == nil {
		rconn, err := redis.Dial("tcp", server)
		if err != nil {
			log.Println("link redis error", err)
			e = tools.CouldNotConnRedisError
		}
		if password != "" {
			if _, err := rconn.Do("AUTH", password); err != nil {
				rconn.Close()
				log.Println("AUTH redis error", err)
				e = tools.RedisAuthError
			}
		}
		if err := rconn.Err(); err == nil {
			MyRedis = &redisConn{
				Conn: rconn,
			}
			log.Println("redis connect success!")
		}
	}
	return MyRedis, e
}
func InitRedis() (myredis *redisConn, e error) {
	if MyRedis == nil {
		MyRedis, e = NewRedis(conf.MyConfig.Redis.Addr+":"+conf.MyConfig.Redis.Port, conf.MyConfig.Redis.Password)
	}
	return
}

func NewPool(server, password string) *redisPool {
	if MyRedisPool == nil {
		MyRedisPool = &redisPool{
			Pool: &redis.Pool{
				MaxIdle:     3,
				IdleTimeout: 240 * time.Second,
				Dial: func() (redis.Conn, error) {
					c, err := redis.Dial("tcp", server)
					if err != nil {
						return nil, err
					}
					if password != "" {
						if _, err := c.Do("AUTH", password); err != nil {
							c.Close()
							return nil, err
						}
					}

					return c, err
				},
				TestOnBorrow: func(c redis.Conn, t time.Time) error {
					_, err := c.Do("PING")
					return err
				},
			},
		}
	}
	return MyRedisPool
}

func DefaultPool() *redisPool {
	if MyRedisPool == nil {
		MyRedisPool = NewPool(conf.MyConfig.Redis.Addr+":"+conf.MyConfig.Redis.Port, conf.MyConfig.Redis.Password)
	}
	return MyRedisPool
}

func (r *redisConn) close() {
	r.Conn.Close()
}

func (r *redisPool) close() {
	r.Pool.Close()
}

func Close() {
	tools.Println(&MyRedis, MyRedis)
	if MyRedis != nil && MyRedis.Conn != nil {
		MyRedis.close()
		log.Println("close redis connect!")
	} else {
		log.Println("didn't found redis connect!")
	}
	if MyRedisPool != nil && MyRedisPool.Pool != nil {
		MyRedisPool.close()
		log.Println("close redis pool!")
	} else {
		log.Println("didn't found redis pool!")
	}
}
