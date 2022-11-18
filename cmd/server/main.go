package main

import (
	"AvitoInternship/config"
	"AvitoInternship/internal/server"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

var configPath string = "config/config.toml"

func main() {
	config := config.CreateServerConfig()
	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		logrus.Fatal(err)
	}
	contextLogger := logrus.WithFields(logrus.Fields{})
	logrus.SetReportCaller(false)
	logrus.SetFormatter(&logrus.TextFormatter{PadLevelText: false, DisableLevelTruncation: false})
	appServer := server.CreateServer(config, contextLogger)

	err = appServer.Start()
	if err != nil {
		panic(err)
	}
}
