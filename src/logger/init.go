package logger

import (
	"os"
	"github.com/BurntSushi/toml"
)

var config struct {
	LogPath       string
	LogFile       string
	EnableConsole bool
	Level         string
	LogLevel      LEVEL
}

func init() {
	loadConfig()
	_, err := os.Stat(config.LogPath)
	if !os.IsExist(err) {
		err = os.MkdirAll(config.LogPath, os.ModePerm)
	}

	SetConsole(config.EnableConsole)
	SetRollingDaily(config.LogPath, config.LogFile)
	SetLevel(config.LogLevel)
}

func loadConfig() {
	_, err := toml.DecodeFile("./conf/logger.conf", &config)
	if err != nil {
		panic(err)
	}
	switch config.Level {
	case "ALL":
		config.LogLevel = ALL
	case "DEBUG":
		config.LogLevel = DEBUG
	case "INFO":
		config.LogLevel = INFO
	case "WARN":
		config.LogLevel = WARN
	case "ERROR":
		config.LogLevel = ERROR
	case "FATAL":
		config.LogLevel = FATAL
	case "OFF":
		config.LogLevel = OFF
	}
}
