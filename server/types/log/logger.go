package log

import (
	"fmt"
	"github.com/dukryung/microservice/server/types/configs"
	"github.com/fatih/color"
	"log"
	"os"
	"strings"
)

type Logger struct {
	prefix string
	enable bool
	level  int
	debug  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	error  *log.Logger
}

func NewLogger(prefix string, config configs.LogConfig) *Logger {
	prefix = strings.Trim(prefix, " ")
	flag := log.Ltime | log.Lshortfile

	l := Logger{
		prefix: prefix,
		enable: config.Enable,
		level:  config.Level,
		debug:  log.New(os.Stdout, color.BlueString("DEBUG "), flag),
		info:   log.New(os.Stdout, color.CyanString("INFO "), flag),
		warn:   log.New(os.Stdout, color.GreenString("WARN "), flag),
		error:  log.New(os.Stdout, color.RedString("ERROR "), flag),
	}
	return &l
}

func (l *Logger) write(logger *log.Logger, logLv int, v ...interface{}) {
	if !l.enable {
		return
	}

	if l.level > logLv {
		return
	}

	prefix := fmt.Sprintf(color.CyanString("x") + "=" + l.prefix + " ")
	logger.Output(3, prefix+fmt.Sprint(v...))
}

func (l *Logger) Debug(v ...interface{}) {
	l.write(l.debug, 2, v)
}

func (l *Logger) Info(v ...interface{}) {
	l.write(l.info, 3, v)
}

func (l *Logger) Warn(v ...interface{}) {
	l.write(l.warn, 4, v)
}

func (l *Logger) Err(v ...interface{}) {
	l.write(l.error, 5, v)
}
