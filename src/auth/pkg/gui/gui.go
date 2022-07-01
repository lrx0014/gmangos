package gui

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"gmangos/src/libs/fifo"
	"gmangos/src/libs/utils"
)

type UI struct {
	log        *logger
	autoScroll bool
	reverseLog bool
}

func New(buf *fifo.Queue) *UI {
	return &UI{
		log:        newLogger(buf),
		autoScroll: true,
	}
}

func (u UI) Draw(wnd string) {
	g.NewMasterWindow(wnd, 400, 200, 0).Run(u.loop)
}

func Update() {
	g.Update()
}

func (u *UI) loop() {
	g.SingleWindow().Layout(
		g.Row(
			g.Button("clear").OnClick(u.clearLogBuf),
			g.Checkbox("reverse", &u.reverseLog),
			g.Checkbox("AutoScroll", &u.autoScroll),
		),

		g.Child().Layout(
			g.Custom(func() {
				for _, msg := range u.readLogBuf() {
					g.Label(msg).Wrapped(true).Build()
				}

				if u.autoScroll {
					defer imgui.SetScrollHereY(1.0)
				}
			}),
		).ID("log_panel"),
	)
}

func (u UI) readLogBuf() []string {
	res := make([]string, 0)
	all := u.log.buf.All()
	for _, data := range all {
		res = append(res, string(data))
	}

	if u.reverseLog {
		utils.ReverseStrings(res)
	}

	return res
}

func (u UI) clearLogBuf() {
	u.log.buf.Reset()
}
