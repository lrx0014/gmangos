package tcp

import (
	"net"
	"sync"
)

type ConnPool struct {
	conn map[int]net.Conn
	lock sync.RWMutex
}

func NewConnPool() *ConnPool {
	return &ConnPool{
		conn: make(map[int]net.Conn),
		lock: sync.RWMutex{},
	}
}

func (c *ConnPool) Put(conn net.Conn) (fd int, err error) {
	fd, err = GetFd(conn)
	if err != nil {
		return
	}

	c.lock.Lock()
	defer c.lock.Unlock()
	c.conn[fd] = conn

	return
}

func (c *ConnPool) Get(fd int) (conn net.Conn) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.conn[fd]
}
