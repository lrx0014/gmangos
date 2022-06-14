package tcp

import (
	"errors"
	"net"
)

var (
	ErrNotTCPConn = errors.New("network: not a tcp connect")
)

func GetFd(conn net.Conn) (fd int, err error) {
	if s, ok := conn.(*net.TCPConn); !ok {
		err = ErrNotTCPConn
		return
	} else {
		f, e := s.File()
		if e != nil {
			err = e
			return
		}

		fd = int(f.Fd())
	}

	return
}
