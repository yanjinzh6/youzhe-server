package conf

import(
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/yanjinzh6/youzhe-server/tools"
)

type Config struct {
	Server serverConfig
	Redis redisConfig
}

type serverConfig struct {
	Addr string
	Port string
	Ports string
	HandlerList []handlerItem
}

type handlerItem struct {
	Action string
	MyFunc string
}

type redisConfig struct {
	Addr string
	Port string
}

var config *Config

func InitConfig(filePath string) (conf *Config) {
	if filePath == "" {
		filePath = tools.DEFAULT_CONFIG_FILE
	}
	if config == nil {
		file, err := tools.LoadFile(filePath)
		tools.Println(err)
		reader := bufio.NewReaderSize(file, tools.DEFAULT_BUFFER_SIZE)
		buf := make([]byte, 0, tools.DEFAULT_BUFFER_SIZE)
		bufs := bytes.NewBuffer(buf)
		for {
			buf, err := reader.ReadBytes('\n')
			tools.Println(buf)
			bufs.Write(buf)
			if err != nil {
				break
			}
		}
		tools.Println(bufs.Len())
		if bufs.Len() == 0 {
			//nothing
			config = &Config{
				Server: serverConfig{Addr: "127.0.0.1", Port: "3333", Ports: "1443", HandlerList: []handlerItem{handlerItem{Action: "/index", MyFunc: "/index/Index"}}},
				Redis: redisConfig{Addr: "127.0.0.1", Port: "6379"},
			}
			tools.Println(config)
			writer := bufio.NewWriterSize(file, tools.DEFAULT_BUFFER_SIZE)
			buf, err = json.Marshal(&config)
			tools.Println(err)
			n, err := writer.Write(buf)
			tools.Println(n, err)
			writer.Flush()
		} else {
			err = json.Unmarshal(bufs.Bytes(), &config)
			tools.Println(err)
		}
	}
	return config
}