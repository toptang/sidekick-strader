package utils

type HttpConfig struct {
	Addr         string `json:"address"`
	Port         int    `json:"port"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
}

var (
	httpConfig *HttpConfig
)

func InitHttp(httpConf HttpConfig) {
	httpConfig = &httpConf
	if httpConfig.Addr == "" ||
		httpConfig.Port == 0 {
		panic("http service config error")
	}
}

func GetHttpAddr() string {
	return httpConfig.Addr
}

func GetHttpPort() int {
	return httpConfig.Port
}

func GetHttpRTimeout() int {
	return httpConfig.ReadTimeout
}

func GetHttpWTimeout() int {
	return httpConfig.WriteTimeout
}
