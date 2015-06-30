package tools

import (
	"errors"
	"fmt"
)

const (
	DEFAULT_CONFIG_FILE = "/conf/config.json"
)

var (
	Debug = true
)

var (
	NilKeyError = errors.New("nil key error")
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