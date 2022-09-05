package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"gmangos/src/auth/pkg/processor"
	"gmangos/src/libs/config"
	"gmangos/src/libs/network/tcp"
	"gmangos/src/libs/welcome"
	"io"
	"os"
)

var (
	configPath  string
	runInDocker bool
)

func main() {
	parseFlags()
	var conf config.Conf

	if "1" == os.Getenv("RUN_IN_DOCKER") {
		// 容器内运行从环境变量获取配置信息
		conf = config.LoadEnv()
		runInDocker = true
	} else {
		// 二进制运行从配置文件获取
		conf = config.LoadConfig(config.ReadFile(configPath))
	}

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
	if runInDocker {
		log.Infof("[gMaNGOS][auth_server] running in docker mode, all config params will be loaded from env")
	}
	log.Infof("[gMaNGOS][auth_server] endpoint: %s:%s", config.C.Server.Host, config.C.Server.Port)

	select {}
}

func parseFlags() {
	path := flag.String("conf", "conf/auth.toml", "指定toml配置文件路径")
	flag.Parse()

	configPath = *path
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
}
