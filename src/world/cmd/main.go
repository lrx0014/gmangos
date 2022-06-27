package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"gmangos/src/libs/config"
	"gmangos/src/libs/constants"
	"gmangos/src/libs/network/tcp"
	"gmangos/src/world/pkg/processor"
)

var (
	configPath string
)

func main() {
	parseFlags()
	conf := config.LoadConfig(config.ReadFile(configPath))

	server, err := tcp.NewServer(conf.Server)
	if err != nil {
		panic(err)
	}

	constants.Welcome()
	log.Infof("[gMaNGOS][world_server] VERSION %s", constants.Version())
	log.Infof("[gMaNGOS][world_server] is running.")
	log.Infof("[gMaNGOS][world_server] endpoint: %s:%s", conf.Server.Addr, conf.Server.Port)

	server.Register(processor.NewWorldProcessor())
	server.Run()
}

func parseFlags() {
	path := flag.String("conf", "config.toml", "指定toml配置文件路径")
	flag.Parse()

	configPath = *path
}
