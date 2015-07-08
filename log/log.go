package log

import (
	"log"
	"github.com/yanjinzh6/youzhe-server/tools"
)

func init() {
	file, err := tools.InitFile(tools.DEFAULT_LOG_FILE)
	if (err == nil) {
		log.SetOutput(file)
		log.SetFlags(log.Flags() | log.Lshortfile)
		log.Println("start log")
	}
}

func Println(arg ...interface{}) {
	log.Println(arg ...)
}