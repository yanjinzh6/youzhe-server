package conf

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/yanjinzh6/youzhe-server/tools"
)

type Config struct {
	Server  serverConfig
	Redis   dbConfig
	Mongodb dbConfig
}

type serverConfig struct {
	Addr        string
	Port        string
	Ports       string
	HandlerList []handlerItem
}

type handlerItem struct {
	Action string
	MyFunc string
}

type dbConfig struct {
	Addr     string
	Port     string
	User     string
	Password string
	DbName   string
}

var MyConfig *Config

func init() {
	MyConfig = InitConfig("")
}

func InitConfig(filePath string) (conf *Config) {
	if filePath == "" {
		filePath = tools.DEFAULT_CONFIG_FILE
	}
	if MyConfig == nil {
		file, err := tools.LoadFile(filePath)
		tools.Println(err)
		reader := bufio.NewReaderSize(file, tools.DEFAULT_BUFFER_SIZE)
		buf := make([]byte, 0, tools.DEFAULT_BUFFER_SIZE)
		bufs := bytes.NewBuffer(buf)
		for {
			buf, err := reader.ReadBytes('\n')
			//tools.Println(buf)
			bufs.Write(buf)
			if err != nil {
				break
			}
		}
		tools.Println(bufs.Len())
		if bufs.Len() == 0 {
			//nothing
			MyConfig = &Config{
				Server:  serverConfig{Addr: "127.0.0.1", Port: "3333", Ports: "1443", HandlerList: []handlerItem{handlerItem{Action: "/index", MyFunc: "/index/Index"}}},
				Redis:   dbConfig{Addr: "127.0.0.1", Port: "6379", User: "", Password: "", DbName: ""},
				Mongodb: dbConfig{Addr: "127.0.0.1", Port: "27017", User: "", Password: "", DbName: "yozh"},
			}
			//tools.Println(MyConfig)
			writer := bufio.NewWriterSize(file, tools.DEFAULT_BUFFER_SIZE)
			buf, err = json.Marshal(&MyConfig)
			tools.Println(err)
			n, err := writer.Write(buf)
			tools.Println(n, err)
			writer.Flush()
		} else {
			err = json.Unmarshal(bufs.Bytes(), &MyConfig)
			tools.Println(err)
		}
	}
	return MyConfig
}
