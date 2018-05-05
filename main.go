package main

import (
	"flag"
	"fmt"
	"sidekick/strader/app"
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
	var app app.Config
	err := config.LoadConfigFromFileV2(&app, *conf)

	//TODO  use errd
	if err != nil {
		panic(fmt.Sprintf("Load configuration error: %v", err))
	}

	//init http service conf
	utils.InitHttp(app.HttpConf)

	//init log
	utils.InitLog(app.LogConf)

	//start service
	if err = server.RunHTTPMux(utils.GetHttpAddr(), utils.GetHttpPort(), http_handler.Rt, utils.GetHttpRTimeout(), utils.GetHttpWTimeout()); err != nil {
		panic(fmt.Sprintf("run tmatric service error: %v", err))
	}
}
