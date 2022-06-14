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
	log.Debugf("[auth][onConnect]")
	return nil
}

func (a AuthProcessor) OnReceive(bytes []byte) error {
	log.Debugf("[auth][onReceive]")
	return nil
}

func (a AuthProcessor) OnClose() {
	log.Debugf("[auth][onClose]")
}
