package main

import (
	"flag"
	"fmt"
	"sidekick/strader/app"
	_ "sidekick/strader/logic/service"
	"sidekick/strader/utils"
	"xframe/cmd"
	"xframe/config"
	"xframe/handler/http_handler"
	"xframe/server"
)

var (
	conf = flag.String("c", "", "configuration file path")
)

func main() {
	//init commandLine
	cmd.ParseCommand()
	cmd.DumpCommand()
	//init configuration
	var appConf app.Config
	err := config.LoadConfigFromFileV2(&appConf, *conf)
	if err != nil {
		panic(fmt.Sprintf("Load configuration error: %v", err))
	}
	//init http service conf
	utils.InitHttp(appConf.HttpConf)
	//init log
	utils.InitLog(appConf.LogConf)
	//init upstream
	utils.InitOkexConfig(appConf.UpstreamConf.OkexConf)
	//start service
	if err = server.RunHTTPMux(utils.GetHttpAddr(), utils.GetHttpPort(), http_handler.Rt, utils.GetHttpRTimeout(), utils.GetHttpWTimeout()); err != nil {
		panic(fmt.Sprintf("run tmatric service error: %v", err))
	}
}
