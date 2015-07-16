package log

import (
	"github.com/yanjinzh6/youzhe-server/tools"
	"log"
)

func init() {
	file, err := tools.InitFile(tools.DEFAULT_LOG_FILE)
	if err == nil {
		log.SetOutput(file)
		log.SetFlags(log.Flags() | log.Lshortfile)
		log.Println("start log")
	}
}

func Println(arg ...interface{}) {
	if tools.SaveLog {
		log.Println(arg...)
	}
}

func Fatalln(arg ...interface{}) {
	if tools.SaveLog {
		log.Fatalln(arg...)
	}
}

func Panicln(arg ...interface{}) {
	if tools.SaveLog {
		log.Panicln(arg...)
	}
}
