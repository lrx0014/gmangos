package tcp

import (
	"bufio"
	"fmt"
	"net"
)

type Server struct {
}

func NewServer(c conf) (s Server) {
	return Server{}
}

func (s Server) OnConnect() {

}

func (s Server) OnReceive(conn net.Conn) {

}

func (s Server) OnClose() {

}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			break
		}

		data := string(buf[:n])
		fmt.Printf("got：%v\n", data)

		_, err = conn.Write([]byte("ok"))
		if err != nil {
			fmt.Printf("write from conn failed, err:%v\n", err)
			break
		}
	}
}
