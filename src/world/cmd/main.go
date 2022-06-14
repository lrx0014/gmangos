package main

import (
	log "github.com/sirupsen/logrus"
	"gmangos/src/libs/constants"
	"gmangos/src/libs/network/tcp"
	"gmangos/src/world/pkg/processor"
)

func main() {
	conf := &tcp.Conf{
		Addr: "127.0.0.1",
		Port: "9999",
	}
	server, err := tcp.NewServer(conf)
	if err != nil {
		panic(err)
	}

	constants.Welcome()
	log.Infof("[gMaNGOS][world_server] VERSION %s", constants.Version())
	log.Infof("[gMaNGOS][world_server] is running.")
	log.Infof("[gMaNGOS][world_server] endpoint: %s:%s", conf.Addr, conf.Port)

	server.Register(processor.NewWorldProcessor())
	server.Run()
}
