package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"gmangos/src/auth/pkg/processor"
	"gmangos/src/libs/config"
	"gmangos/src/libs/constants"
	"gmangos/src/libs/network/tcp"
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

	server.Register(processor.NewAuthProcessor(conf))
	go server.Run()

	constants.Welcome()
	log.Infof("[gMaNGOS][auth_server] VERSION %s", constants.Version())
	log.Infof("[gMaNGOS][auth_server] is running.")
	log.Infof("[gMaNGOS][auth_server] endpoint: %s:%s", conf.Server.Host, conf.Server.Port)

	select {}
}

func parseFlags() {
	path := flag.String("conf", "config.toml", "指定toml配置文件路径")
	flag.Parse()

	configPath = *path
}
