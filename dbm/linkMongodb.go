package dbm

import (
	"github.com/yanjinzh6/youzhe-server/conf"
	"github.com/yanjinzh6/youzhe-server/log"
	"gopkg.in/mgo.v2"
)

type mongodbSession struct {
	Msession *mgo.Session
}

var MyMongodb *mongodbSession

func InitMongodb() (mdb *mongodbSession, e error) {
	if MyMongodb == nil {
		if ndb, err := NewMongodb(conf.MyConfig.Mongodb.Addr + ":" + conf.MyConfig.Mongodb.Port); err == nil {
			MyMongodb = ndb
		}
	}
	return MyMongodb, e
}

func NewMongodb(url string) (mdb *mongodbSession, e error) {
	if MyMongodb == nil {
		session, err := mgo.Dial(url)
		if err != nil {
			log.Println("connect to mongodb server error", err)
		} else {
			MyMongodb = &mongodbSession{
				Msession: session,
			}
			log.Println("connect to mongodb server success")
		}
	}
	return MyMongodb, e
}

func (m *mongodbSession) close() {
	m.Msession.Close()
}

func Close() {
	if MyMongodb != nil && MyMongodb.Msession != nil {
		MyMongodb.close()
		log.Println("close mongodb session success")
	} else {
		log.Println("could not found mongodb session")
	}
}
