package processor

import (
	log "github.com/sirupsen/logrus"
	"gmangos/src/libs/network/tcp"
)

type AuthProcessor struct {
}

var _ tcp.Processor = new(AuthProcessor)

func NewAuthProcessor() *AuthProcessor {
	return &AuthProcessor{}
}

func (a AuthProcessor) OnConnect() error {
	log.Info("[auth][onConnect]")
	return nil
}

func (a AuthProcessor) OnReceive(bytes []byte) error {
	log.Infof("[auth][onReceive] %+v", bytes)
	return nil
}

func (a AuthProcessor) OnClose() {
	log.Info("[auth][onClose]")
}
