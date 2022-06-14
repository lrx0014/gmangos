package tcp

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"net"
)

type Server struct {
	conf      *Conf
	pool      *ConnPool
	processor Processor
}

type Processor interface {
	OnConnect() error
	OnReceive([]byte) error
	OnClose()
}

func NewServer(c *Conf) (s *Server, err error) {
	err = c.parseConfig()
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
	listen, err := net.Listen("tcp", s.conf.address())
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
		e = s.processor.OnConnect()
		if e != nil {
			log.Errorf("[onConnect] err: %s", e.Error())
			continue
		}

		go s.process(conn)
	}
}

func (s *Server) Shutdown() {

}

func (s *Server) process(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf []byte
		n, err := reader.Read(buf)
		if err != nil {
			log.Errorf("read from conn failed, err: %s", err.Error())
			break
		}

		data := buf[:n]
		e := s.processor.OnReceive(data)
		if e != nil {
			log.Errorf("[onReceive] err: %s", e.Error())
			break
		}
	}
}

type DefaultProcessor struct {
}

func (d *DefaultProcessor) OnConnect() error {
	return nil
}

func (d *DefaultProcessor) OnReceive(data []byte) error {
	return nil
}

func (d *DefaultProcessor) OnClose() {

}
