package conf

type Config struct {
	AppName   string
	RpcServer *RpcServer
}

type RpcServer struct {
	Name     string
	TTL      int
	Interval int
}
