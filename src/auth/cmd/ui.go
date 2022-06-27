package main

import (
	"fmt"
	g "github.com/AllenDang/giu"
	log "github.com/sirupsen/logrus"
)

func loop() {
	g.SingleWindow().Layout(
		g.Labelf(content),
	)
}

var content string

type uiLogger struct {
}

var _ log.Hook = new(uiLogger)

func (u uiLogger) Levels() []log.Level {
	return log.AllLevels
}

func (u uiLogger) Fire(entry *log.Entry) error {
	content += fmt.Sprintf("%s\n", entry.Message)
	return nil
}
