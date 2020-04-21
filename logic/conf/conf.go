package conf

import (
	"fmt"
	"os"

	"github.com/freezeChen/studio-library/database/mysql"
	"github.com/freezeChen/studio-library/redis"
	"github.com/freezeChen/studio-library/util"
	"github.com/freezeChen/studio-library/zlog"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
)

var (
	fileName string
	conf     *Config
)

func init() {
	fileName = os.Getenv("config")
	if len(fileName) == 0 {
		fileName = "conf.yaml"
	}
}

func Init() error {
	conf = new(Config)
	baseConfig, err := config.NewConfig(config.WithSource(file.NewSource(file.WithPath(fmt.Sprintf("%s/%s", util.GetCurrentDirectory(), fileName)))))
	if err != nil {
		return err
	}
	err = baseConfig.Scan(conf)
	return err
}

func GetConf() *Config {
	return conf
}

type Config struct {
	AppName   string        `json:"appName"`
	RpcServer *RpcServer    `json:"rpcServer"`
	Mysql     *mysql.Config `json:"mysql"`
	Redis     *redis.Config `json:"redis"`
	Log       *zlog.Config  `json:"log"`
}

type RpcServer struct {
	Name     string
	TTL      int
	Interval int
}
