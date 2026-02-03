package unilog

import (
	"os"
	"strings"

	"github.com/DRVTiny/lib-tools/tools"
	"github.com/DRVTiny/lib-unilog/unilog/stdlog"
)

type UniLogger interface {
	Fatalf(fstr string, args ...any)
	Fatal(args ...any)
	Errorf(fstr string, args ...any)
	Error(args ...any)
	Warnf(fstr string, args ...any)
	Warn(args ...any)
	Infof(fstr string, args ...any)
	Info(args ...any)
	Debugf(fstr string, args ...any)
	Debug(args ...any)
	Print(args ...any)
	Printf(fstr string, args ...any)
	Println(args ...any)
}

var uniLog UniLogger

func Anys2LogString(args []any) string {
	return strings.Join(tools.Anys2Strings(args), " ")
}

func SetGlobal(l UniLogger) {
	if l == nil {
		panic("Could not use nil interface as global logger. First think, what you do. No log - no work")
	}
	uniLog = l
}

func GetGlobal() UniLogger {
	if uniLog == nil {
		u4l, err := stdlog.NewUniStdLog(os.Stdout, "debug")
		if err != nil {
			panic("this can never happen: NewUniStdLog returns error while using correct log level")
		}

		SetGlobal(u4l)
	}

	return uniLog
}

func Fatalf(fstr string, args ...any) {
	GetGlobal().Fatalf(fstr, args...)
}

func Fatal(args ...any) {
	GetGlobal().Fatal(args...)
}

func Errorf(fstr string, args ...any) {
	GetGlobal().Errorf(fstr, args...)
}

func Error(args ...any) {
	GetGlobal().Error(args...)
}

func Warnf(fstr string, args ...any) {
	GetGlobal().Warnf(fstr, args...)
}

func Warn(args ...any) {
	GetGlobal().Warn(args...)
}

func Infof(fstr string, args ...any) {
	GetGlobal().Infof(fstr, args...)
}

func Info(args ...any) {
	GetGlobal().Info(args...)
}

func Debugf(fstr string, args ...any) {
	GetGlobal().Debugf(fstr, args...)
}

func Debug(args ...any) {
	GetGlobal().Debug(args...)
}

func Printf(fstr string, args ...any) {
	GetGlobal().Debugf(fstr, args...)
}

func Print(args ...any) {
	GetGlobal().Debug(args...)
}

func Println(args ...any) {
	GetGlobal().Debug(args...)
}
