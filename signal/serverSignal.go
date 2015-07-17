package signal

import (
	"github.com/yanjinzh6/youzhe-server/db"
	"github.com/yanjinzh6/youzhe-server/dbm"
	"github.com/yanjinzh6/youzhe-server/log"
	"os"
	"os/signal"
	"syscall"
)

type signalHandler func(s os.Signal, arg interface{})

type signalMap struct {
	sMap map[os.Signal]signalHandler
}

var mySignalMap *signalMap

func init() {
	mySignalMap = New()
	go signalNotify(mySignalMap)
}

func New() (s *signalMap) {
	if mySignalMap == nil {
		mySignalMap = &signalMap{
			sMap: make(map[os.Signal]signalHandler),
		}
	}
	mySignalMap.register(syscall.SIGINT, myHandler)
	mySignalMap.register(syscall.SIGTERM, myHandler)
	return mySignalMap
}

func (sig *signalMap) register(s os.Signal, handler signalHandler) {
	/*if _, found := sig.sMap[s]; !found {
		sig.sMap[s] = handler
	}*/
	sig.sMap[s] = handler
}

func (sig *signalMap) handle(s os.Signal, arg interface{}) {
	if handler, found := sig.sMap[s]; found {
		handler(s, arg)
	} else {
		handler(s, arg)
	}
}

func myHandler(s os.Signal, arg interface{}) {
	log.Println("handle signal", s)
	dbm.Close()
	db.Close()
	os.Exit(9)
}

func signalNotify(mySigMap *signalMap) {
	c := make(chan os.Signal)
	var sigs []os.Signal
	for s, _ := range mySigMap.sMap {
		sigs = append(sigs, s)
	}

	signal.Notify(c)

	for {
		mySigMap.handle(<-c, nil)
	}
}
