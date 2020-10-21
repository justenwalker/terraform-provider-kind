package provider

import (
	"fmt"
	stdlog "log"
	"sigs.k8s.io/kind/pkg/log"
	"strings"
)

var _ log.Logger = logger{}
var _ log.InfoLogger = logger{}

type logger struct {
	Logger   *stdlog.Logger
	MaxLevel log.Level
	CurLevel log.Level
}

func (l logger) Info(message string) {
	if !l.Enabled() {
		return
	}
	if l.CurLevel > 0 {
		l.Logger.Println(fmt.Sprintf("[INFO.%d]", l.CurLevel), message)
	}
	l.Logger.Println("[INFO]", message)
}

func (l logger) Infof(format string, args ...interface{}) {
	if !l.Enabled() {
		return
	}
	if l.CurLevel > 0 {

		l.Logger.Println(fmt.Sprintf("[INFO.%d]", l.CurLevel), fmt.Sprintf(format, args...))
	}

	l.Logger.Println("[INFO]", fmt.Sprintf(format, args...))
}

func (l logger) Enabled() bool {
	return l.CurLevel <= l.MaxLevel
}

func (l logger) Warn(message string) {
	l.log("[WARN]", message)
}

func (l logger) Warnf(format string, args ...interface{}) {
	l.logf("[WARN]", format, args...)
}

func (l logger) Error(message string) {
	l.log("[ERROR]", message)
}

func (l logger) Errorf(format string, args ...interface{}) {
	l.logf("[ERROR]", format, args...)
}

func (l logger) V(level log.Level) log.InfoLogger {
	return logger{
		Logger:   l.Logger,
		MaxLevel: l.MaxLevel,
		CurLevel: l.CurLevel + level,
	}
}

func (l logger) log(prefix string, message string) {
	l.Logger.Println(prefix, strings.TrimSpace(message))
}

func (l logger) logf(prefix string, format string, args ...interface{}) {
	l.Logger.Println(prefix, strings.TrimSpace(fmt.Sprintf(format, args...)))
}
