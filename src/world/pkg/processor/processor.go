package processor

import (
	log "github.com/sirupsen/logrus"
	"gmangos/src/libs/network/tcp"
)

type WorldProcessor struct {
}

var _ tcp.Processor = new(WorldProcessor)

func NewWorldProcessor() *WorldProcessor {
	return &WorldProcessor{}
}

func (a WorldProcessor) OnConnect(fd int) error {
	log.Infof("[world][onConnect][client-%d]", fd)
	return nil
}

func (a WorldProcessor) OnReceive(fd int, bytes []byte) (resp []byte, err error) {
	log.Infof("[world][onReceive][client-%d] %+v", fd, bytes)
	return []byte("hi"), nil
}

func (a WorldProcessor) OnClose(fd int) {
	log.Infof("[world][onClose][client-%d]", fd)
}
