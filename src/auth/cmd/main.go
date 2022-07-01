package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"gmangos/src/auth/pkg/gui"
	"gmangos/src/auth/pkg/processor"
	"gmangos/src/libs/config"
	"gmangos/src/libs/fifo"
	"gmangos/src/libs/network/tcp"
	"gmangos/src/libs/welcome"
	"io"
	"os"
)

type logHook struct {
	loggerBuf *fifo.Queue
}

var _ log.Hook = new(logHook)

func (l logHook) Levels() []log.Level {
	return log.AllLevels
}

func (l logHook) Fire(entry *log.Entry) error {
	l.loggerBuf.Push([]byte(entry.Message + "\n"))
	return nil
}

var (
	configPath string
	runGui     bool

	loggerBuf *fifo.Queue
)

func main() {
	parseFlags()
	conf := config.LoadConfig(config.ReadFile(configPath))

	initEnv()

	server, err := tcp.NewServer(conf.Server)
	if err != nil {
		panic(err)
	}

	server.Register(processor.NewAuthProcessor(conf))
	go server.Run()

	welcome.Welcome()
	log.Infof("[gMaNGOS][auth_server] VERSION %s", welcome.Version())
	log.Infof("[gMaNGOS][auth_server] is running.")
	log.Infof("[gMaNGOS][auth_server] endpoint: %s:%s", config.C.Server.Host, config.C.Server.Port)

	if runGui {
		runInGui()
	} else {
		select {}
	}
}

func parseFlags() {
	path := flag.String("conf", "conf/auth.toml", "指定toml配置文件路径")
	useGui := flag.Bool("gui", false, "以gui模式运行")
	flag.Parse()

	configPath = *path
	runGui = *useGui
}

func initEnv() {
	// init logger
	// log to file
	logPath := config.C.Server.LogPath
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	writers := []io.Writer{
		file,
		os.Stdout, // log to stdout
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	log.SetOutput(fileAndStdoutWriter)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		TimestampFormat:           "2006-01-02 15:04:05",
	})
	// log to memory buffer
	loggerBuf = fifo.New(config.C.Server.LogCacheSize)
	log.AddHook(logHook{loggerBuf: loggerBuf})
}

func runInGui() {
	log.Infof("[gMaNGOS][auth_server] running in GUI mode")
	ui := gui.New(loggerBuf)
	ui.Draw("gMaNGOS auth server")
}
