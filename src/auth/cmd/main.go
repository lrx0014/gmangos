package main

import (
	"flag"
	g "github.com/AllenDang/giu"
	log "github.com/sirupsen/logrus"
	"gmangos/src/auth/pkg/processor"
	"gmangos/src/libs/config"
	"gmangos/src/libs/network/tcp"
	"gmangos/src/libs/welcome"
)

var (
	configPath string
	runGui     bool
)

func main() {
	parseFlags()
	conf := config.LoadConfig(config.ReadFile(configPath))

	if runGui {
		log.AddHook(&uiLogger{})
	}

	server, err := tcp.NewServer(conf.Server)
	if err != nil {
		panic(err)
	}

	server.Register(processor.NewAuthProcessor(conf))
	go server.Run()

	welcome.Welcome()
	log.Infof("[gMaNGOS][auth_server] VERSION %s", welcome.Version())
	log.Infof("[gMaNGOS][auth_server] is running.")
	log.Infof("[gMaNGOS][auth_server] endpoint: %s:%s", config.C.Server.Host, config.C.Server.Port)

	if runGui {
		runInGui()
	} else {
		select {}
	}
}

func parseFlags() {
	path := flag.String("conf", "config.toml", "指定toml配置文件路径")
	useGui := flag.Bool("gui", false, "以gui模式运行")
	flag.Parse()

	configPath = *path
	runGui = *useGui
}

func runInGui() {
	log.Infof("[gMaNGOS][auth_server] running in GUI mode")
	wnd := g.NewMasterWindow("gMaNGOS auth server", 400, 200, 0)
	wnd.Run(loop)
}
