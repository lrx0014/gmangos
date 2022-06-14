package main

import (
	log "github.com/sirupsen/logrus"
	"gmangos/src/auth/pkg/processor"
	"gmangos/src/libs/constants"
	"gmangos/src/libs/network/tcp"
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
	log.Info("[gMaNGOS][auth_server] is running.")
	log.Infof("[gMaNGOS][auth_server] Endpoint: %s:%s", conf.Addr, conf.Port)

	server.Register(processor.NewAuthProcessor())
	server.Run()
}
