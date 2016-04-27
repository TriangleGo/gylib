package logger

import (
	"strconv"
	"testing"
	"time"
)

func writeLog(i int) {
	Debug("Debug>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
	Debugf(">>>>>>>>>>>>>>>>>>>>>>%s", strconv.Itoa(i))
	Info("Info>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
	Warn("Warn>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
	Error("Error>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
	Fatal("Fatal>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
}

func Test(t *testing.T) {
	InitLogger("./logger_config.conf")

	for i := 3; i > 0; i-- {
		go writeLog(i)
		time.Sleep(1000 * time.Millisecond)
	}
	time.Sleep(3 * time.Second)
}
