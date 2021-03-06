package utils

import "xframe/log"

type LogConfig struct {
	LogLevel   string `json:"log_level"`
	StackTrace bool   `json:"stack_trace"`
}

//default: using stdout
func InitLog(logConf LogConfig) {
	if logConf.LogLevel == "" {
		panic("log configuration error")
	}
	log.InitLogger("", "", "", 0, logConf.LogLevel, "stdout")
	log.EnableLogDepth(logConf.StackTrace)
}
