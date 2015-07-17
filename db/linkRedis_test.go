package db

import (
	"github.com/yanjinzh6/youzhe-server/tools"
	"testing"
)

func TestInitRedis(t *testing.T) {
	tools.Println(InitRedis())
	Close()
}
