package app

import "sidekick/strader/utils"

type Config struct {
	HttpConf     utils.HttpConfig `json:'http"`
	LogConf      utils.LogConfig  `json:"log"`
	UpstreamConf struct {
		OkexConf utils.OkexConfig `json:"okex"`
	} `json:"upstream"`
}
