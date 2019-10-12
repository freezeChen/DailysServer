package conf

import (
	"fmt"
	"os"

	"github.com/freezeChen/studio-library/util"
	"github.com/freezeChen/studio-library/zlog"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
)

var (
	fileName string
	Conf     *Config
)

func init() {
	fileName = os.Getenv("config")
	if len(fileName) == 0 {
		fileName = "conf.yaml"
	}
}

func Init() error {
	Conf = new(Config)
	baseConfig := config.NewConfig(config.WithSource(file.NewSource(file.WithPath(fmt.Sprintf("%s/%s", util.GetCurrentDirectory(), fileName)))))
	err := baseConfig.Scan(Conf)
	return err
}

type Config struct {
	AppName   string           `json:"appName"`
	RpcServer *RpcServer       `json:"rpcServer"`
	WebSocket *WebSocketConfig `json:"webSocket"`
	TCPConfig *TCPConfig       `json:"tcpConfig"`
	Log       *zlog.Config     `json:"log"`
}

type RpcServer struct {
	Name     string
	TTL      int
	Interval int
}

type WebSocketConfig struct {
	Addr string
}

type TCPConfig struct {
	Port           string
	TCPKeepalive   bool
	TCPReadBuffer  int
	TCPWriteBuffer int
}
