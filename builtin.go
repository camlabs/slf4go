package slf4go

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	config "github.com/dynamicgo/go-config"
	"github.com/fatih/color"
)

var fatalp = color.New(color.FgRed).PrintFunc()
var fatalf = color.New(color.FgRed).PrintfFunc()

var errorp = color.New(color.FgRed).PrintFunc()
var errorf = color.New(color.FgRed).PrintfFunc()

var warnp = color.New(color.FgYellow).PrintFunc()
var warnf = color.New(color.FgYellow).PrintfFunc()

var infop = color.New(color.FgWhite).PrintFunc()
var infof = color.New(color.FgWhite).PrintfFunc()

var debugp = color.New(color.FgCyan).PrintFunc()
var debugf = color.New(color.FgCyan).PrintfFunc()

var tracep = color.New(color.FgBlue).PrintFunc()
var tracef = color.New(color.FgBlue).PrintfFunc()

type colorConsole struct {
	messages chan func()
}

func newColorConsole() LoggerFactory {
	console := &colorConsole{
		messages: make(chan func(), 1000),
	}

	return console
}

func (console *colorConsole) runLoop() {
	for f := range console.messages {
		f()
	}
}

func (console *colorConsole) GetLogger(name string) Logger {
	return &colorConsoleLogger{name: name, messages: console.messages}
}

type colorConsoleLogger struct {
	name     string
	messages chan func()
}

func (logger *colorConsoleLogger) GetName() string {
	return logger.name
}

func source() string {
	_, filename, line, _ := runtime.Caller(3)

	return fmt.Sprintf("%s:%d", filepath.Base(filename), line)
}

func (logger *colorConsoleLogger) Trace(args ...interface{}) {

	logger.messages <- func() {
		tracef("[%s][%s][%s] TRACE ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
		tracep(args...)
		tracep("\n")
	}

}

func (logger *colorConsoleLogger) TraceF(format string, args ...interface{}) {
	logger.messages <- func() {
		tracef("[%s][%s][%s] TRACE ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
		tracef(format, args...)
		tracep("\n")
	}
}

func (logger *colorConsoleLogger) Debug(args ...interface{}) {
	logger.messages <- func() {
		debugf("[%s][%s][%s] DEBUG ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
		debugp(args...)
		debugp("\n")
	}
}

func (logger *colorConsoleLogger) DebugF(format string, args ...interface{}) {
	logger.messages <- func() {
		debugf("[%s][%s][%s] DEBUG ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
		debugf(format, args...)
		debugp("\n")
	}
}

func (logger *colorConsoleLogger) Info(args ...interface{}) {
	logger.messages <- func() {
		infof("[%s][%s][%s] INFO  ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
		infop(args...)
		infop("\n")
	}
}

func (logger *colorConsoleLogger) InfoF(format string, args ...interface{}) {
	logger.messages <- func() {
		infof("[%s][%s][%s] INFO  ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
		infof(format, args...)
		infop("\n")
	}
}

func (logger *colorConsoleLogger) Warn(args ...interface{}) {
	logger.messages <- func() {
		warnf("[%s][%s][%s] WARN  ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
		warnp(args...)
		warnp("\n")
	}
}

func (logger *colorConsoleLogger) WarnF(format string, args ...interface{}) {
	logger.messages <- func() {
		warnf("[%s][%s][%s] WARN  ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
		warnf(format, args...)
		warnp("\n")
	}
}

func (logger *colorConsoleLogger) Error(args ...interface{}) {
	logger.messages <- func() {
		errorf("[%s][%s][%s] ERROR ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
		errorp(args...)
		errorp("\n")
	}
}

func (logger *colorConsoleLogger) ErrorF(format string, args ...interface{}) {
	logger.messages <- func() {
		errorf("[%s][%s][%s] ERROR ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
		errorf(format, args...)
		errorp("\n")
	}
}

func (logger *colorConsoleLogger) Fatal(args ...interface{}) {
	logger.messages <- func() {
		fatalf("[%s][%s][%s] FATAL ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
		fatalp(args...)
		fatalp("\n")
	}
}

func (logger *colorConsoleLogger) FatalF(format string, args ...interface{}) {
	logger.messages <- func() {
		fatalf("[%s][%s][%s] FATAL ", time.Now().Format("2006-01-02 15:04:05"), logger.name, source())
		fatalf(format, args...)
		fatalp("\n")
	}
}

func init() {
	println("[slf4go] register console backend")
	RegisterBackend("console", func(config config.Config) (LoggerFactory, error) {
		return newColorConsole(), nil
	})
}
