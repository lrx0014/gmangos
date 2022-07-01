package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"gmangos/src/libs/config"
	"gmangos/src/libs/network/tcp"
	"gmangos/src/libs/welcome"
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

	server.Register(processor.NewWorldProcessor(conf))
	go server.Run()

	welcome.Welcome()
	log.Infof("[gMaNGOS][world_server] VERSION %s", welcome.Version())
	log.Infof("[gMaNGOS][world_server] is running.")
	log.Infof("[gMaNGOS][world_server] endpoint: %s:%s", conf.Server.Host, conf.Server.Port)

	select {}
}

func parseFlags() {
	path := flag.String("conf", "conf/world.toml", "指定toml配置文件路径")
	flag.Parse()

	configPath = *path
}
