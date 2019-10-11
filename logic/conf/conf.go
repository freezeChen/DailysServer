package conf

import (
	"fmt"
	"os"

	"github.com/freezeChen/studio-library/database/mysql"
	"github.com/freezeChen/studio-library/redis"
	"github.com/freezeChen/studio-library/util"
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
	AppName   string
	RpcServer *RpcServer
	Mysql     *mysql.Config
	Redis     *redis.Config
}

type RpcServer struct {
	Name     string
	TTL      int
	Interval int
}
