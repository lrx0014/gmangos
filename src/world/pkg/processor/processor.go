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

func (a WorldProcessor) OnConnect() error {
	log.Debugf("[world][onConnect]")
	return nil
}

func (a WorldProcessor) OnReceive(bytes []byte) error {
	log.Debugf("[world][onReceive]")
	return nil
}

func (a WorldProcessor) OnClose() {
	log.Debugf("[world][onClose]")
}
