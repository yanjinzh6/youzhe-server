package db

import (
	"gopkg.in/mgo.v2"
)

type mongodbSession struct {
	Msession mgo.Session
}

var MyMongodb mongodbSession

func InitMongodb() {

}
