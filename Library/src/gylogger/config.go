package logger

import (
	"os"
	"github.com/stvp/go-toml-config"
)

var (
	logPath string
	logFile string
	enableConsole bool
	level LEVEL
)

func InitLogger(path string) {
	Infof("Init logger with given config file %s", path)
	loadConfig(path)
	_, err := os.Stat(logPath)
	if !os.IsExist(err) {
		err = os.MkdirAll(logPath, os.ModePerm)
	}

	SetConsole(enableConsole)
	SetRollingDaily(logPath, logFile)
	SetLevel(level)

	if err != nil {
		Fatal("Init log path error", err)
	} else {
		Info("Log path initialised.")
	}
}

func loadConfig(path string) {
	logConfig := config.NewConfigSet("logConfig", config.ExitOnError)
	logConfig.StringVar(&logPath, "log_path", "./default_logs/")
	logConfig.StringVar(&logFile, "log_file", "default.log")
	var tempLevelStr string
	logConfig.StringVar(&tempLevelStr, "log_level", "INFO")
	logConfig.BoolVar(&enableConsole, "enable_console", false)

	err := logConfig.Parse(path)
	if err != nil {
		Warnf("load logger config error, %v", err)
	} else {
		Infof("loaded logger config level = %d", level)
	}

	switch tempLevelStr {
	case "ALL":
		level = ALL
	case "DEBUG":
		level = DEBUG
	case "INFO":
		level = INFO
	case "WARN":
		level = WARN
	case "ERROR":
		level = ERROR
	case "FATAL":
		level = FATAL
	case "OFF":
		level = OFF
	}

}