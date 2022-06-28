package gui

import (
	"gmangos/src/libs/fifo"
)

type logger struct {
	buf *fifo.Queue
}

func newLogger(buf *fifo.Queue) *logger {
	return &logger{buf: buf}
}
