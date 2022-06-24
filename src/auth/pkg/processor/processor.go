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

func (a AuthProcessor) OnConnect(fd int) error {
	log.Infof("[auth][onConnect][client-%d]", fd)
	return nil
}

func (a AuthProcessor) OnReceive(fd int, bytes []byte) (resp []byte, err error) {
	log.Infof("[auth][onReceive][client-%d] %+v", fd, bytes)
	return []byte("hi"), nil
}

func (a AuthProcessor) OnClose(fd int) {
	log.Infof("[auth][onClose][client-%d]", fd)
}
