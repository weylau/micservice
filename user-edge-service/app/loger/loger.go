package loger

// Package log 基础日志组件
import (
	"github.com/k0kubun/pp"
	"github.com/mattn/go-isatty"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"user-edge-service/app/config"
	"user-edge-service/app/helper"
	"os"
	"time"
)

func init() {
	setLevel()
	//initPP()
}

var Loger *logrus.Logger

func Default() *logrus.Logger {
	if Loger != nil {
		return Loger
	}
	appDir := helper.GetAppDir()
	today := time.Now().Format("2006-01-02")
	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  appDir + "/log/info-" + today + ".log",
		logrus.ErrorLevel: appDir + "/log/error-" + today + ".log",
	}

	Loger = logrus.New()

	Loger.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))
	return Loger
}

func initPP() {
	out := os.Stdout
	pp.SetDefaultOutput(out)

	if !isatty.IsTerminal(out.Fd()) {
		pp.ColoringEnabled = false
	}
}

var levels = map[string]logrus.Level{
	"panic": logrus.PanicLevel,
	"fatal": logrus.FatalLevel,
	"error": logrus.ErrorLevel,
	"warn":  logrus.WarnLevel,
	"info":  logrus.InfoLevel,
	"debug": logrus.DebugLevel,
}

func setLevel() {
	levelConf := config.Configs.LogLevel

	if levelConf == "" {
		levelConf = "info"
	}

	if level, ok := levels[levelConf]; ok {
		logrus.SetLevel(level)
	} else {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

// PP 类似 PHP 的 var_dump
func PP(args ...interface{}) {
	pp.Println(args...)
}
