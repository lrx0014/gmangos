package gui

import (
	g "github.com/AllenDang/giu"
	"gmangos/src/libs/fifo"
)

type UI struct {
	log *logger
}

func New(buf *fifo.Queue) *UI {
	return &UI{
		log: newLogger(buf),
	}
}

func (u UI) Draw(wnd string) {
	g.NewMasterWindow(wnd, 400, 200, 0).Run(u.loop)
}

func (u UI) loop() {
	g.SingleWindow().Layout(
		g.Labelf(u.readLogBuf()),
	)
}

func (u UI) readLogBuf() string {
	res := ""
	all := u.log.buf.All()
	for _, data := range all {
		res += string(data)
	}

	return res
}
