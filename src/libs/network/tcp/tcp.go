package tcp

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"gmangos/src/libs/config"
	"io"
	"net"
)

type Server struct {
	conf      *config.ServerConf
	pool      *ConnPool
	processor Processor
}

type Processor interface {
	OnConnect(fd int) (err error)
	OnReceive(fd int, req []byte) (resp []byte, err error)
	OnClose(fd int)
}

func NewServer(c *config.ServerConf) (s *Server, err error) {
	err = c.ParseConfig()
	if err != nil {
		return
	}

	s = &Server{
		conf:      c,
		pool:      NewConnPool(),
		processor: &DefaultProcessor{},
	}

	return
}

func (s *Server) Register(p Processor) {
	s.processor = p
}

func (s *Server) Run() {
	listen, err := net.Listen("tcp", s.conf.Address())
	if err != nil {
		log.Errorf("[listen] err: %s", err.Error())
		return
	}

	for {
		conn, e := listen.Accept()
		if e != nil {
			log.Errorf("[accept] err: %s", e.Error())
			continue
		}
		fd, e := s.pool.Put(conn)
		if e != nil {
			log.Errorf("[pool][put] err: %s", e.Error())
			continue
		}
		e = s.processor.OnConnect(fd)
		if e != nil {
			log.Errorf("[onConnect] err: %s", e.Error())
			continue
		}

		go s.process(fd, conn)
	}
}

func (s *Server) Shutdown() {

}

func (s *Server) process(fd int, conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			s.processor.OnClose(fd)
			break
		}
		if err != nil {
			log.Errorf("[onReceive] read tcp err: %s", err.Error())
			break
		}

		data := buf[:n]
		if len(data) > 0 {
			resp, e := s.processor.OnReceive(fd, data)
			if e != nil {
				log.Errorf("[onReceive] process data err: %s", e.Error())
				break
			}
			_, e = conn.Write(resp)
			if e != nil {
				log.Errorf("[onReceive] write tcp err: %s", e.Error())
				break
			}
		}
	}
}

type DefaultProcessor struct {
}

func (d *DefaultProcessor) OnConnect(fd int) error {
	return nil
}

func (d *DefaultProcessor) OnReceive(fd int, data []byte) (resp []byte, err error) {
	return nil, nil
}

func (d *DefaultProcessor) OnClose(fd int) {

}
