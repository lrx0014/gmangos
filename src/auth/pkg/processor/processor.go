package processor

import (
	log "github.com/sirupsen/logrus"
	"gmangos/src/dao"
	"gmangos/src/libs/config"
	"gmangos/src/libs/network/tcp"
	"time"
)

type AuthProcessor struct {
	dao *dao.Dao
}

var _ tcp.Processor = new(AuthProcessor)

func NewAuthProcessor(c config.Conf) *AuthProcessor {
	return &AuthProcessor{
		dao: dao.New(c),
	}
}

func (a AuthProcessor) OnConnect(fd int) error {
	log.Infof("[auth][onConnect][client-%d]", fd)
	err := a.dao.SaveConnState(fd, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (a AuthProcessor) OnReceive(fd int, bytes []byte) (resp []byte, err error) {
	log.Infof("[auth][onReceive][client-%d] %+v", fd, bytes)
	return []byte("hi"), nil
}

func (a AuthProcessor) OnClose(fd int) {
	log.Infof("[auth][onClose][client-%d]", fd)
}
