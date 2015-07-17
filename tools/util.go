package tools

import (
	"errors"
	"fmt"
	"os"
	"path"
)

const (
	DEFAULT_CONFIG_FILE = "conf/config.json"
	DEFAULT_LOG_FILE    = "log/debug.log"
	DEFAULT_BUFFER_SIZE = 2048
)

var (
	Debug   = true
	SaveLog = true
)

var (
	NilKeyError            = errors.New("nil key error")
	KeyNotNilError         = errors.New("key-value not nil error")
	FuncNotFoundError      = errors.New("can not found the func error")
	ActionNotFoundError    = errors.New("can not found the action error")
	ActionParamError       = errors.New("now need the param like /url/url2")
	CouldNotConnRedisError = errors.New("could not connect to redis server")
	RedisAuthError         = errors.New("could not auth the redis server")
)

func Printf(format string, a ...interface{}) (n int, err error) {
	if Debug {
		return fmt.Printf(format, a...)
	} else {
		return 0, nil
	}
}

func Println(a ...interface{}) (n int, err error) {
	if Debug {
		return fmt.Println(a...)
	} else {
		return 0, nil
	}
}

func LoadFile(filePath string) (file *os.File, err error) {
	file, err = os.OpenFile(filePath, os.O_RDWR, 0x0666)
	if err != nil {
		if os.IsExist(err) {
			Println(err)
		} else {
			err = os.MkdirAll(path.Dir(filePath), 0x0666)
			if err != nil {
				Println(err)
			} else {
				file, err = os.Create(filePath)
				if err != nil {
					Println(err)
				}
			}
		}
	}
	return
}

func InitFile(filePath string) (file *os.File, err error) {
	file, err = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0x0666)
	if err != nil {
		if os.IsExist(err) {
			Println(err)
		} else {
			err = os.MkdirAll(path.Dir(filePath), 0x0666)
			if err != nil {
				Println(err)
			} else {
				file, err = os.Create(filePath)
				if err != nil {
					Println(err)
				}
			}
		}
	}
	return
}
