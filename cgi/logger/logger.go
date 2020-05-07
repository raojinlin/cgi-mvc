package logger

import (
	"bytes"
	"io"
	"log"
	"os"
)

const (
	LogInfo = 0
	LogError
	LogDebug
	LogWarn
)

type Logger struct {
	debug bool
	name string
	logLevel int
	errorWriter io.Writer
	log.Logger
}

func (l *Logger) Log(msg string) {
	if l.logLevel == LogInfo {
		l.Info(msg)
	} else if l.logLevel == LogDebug {
		l.Debug(msg)
	} else if l.logLevel == LogError {
		l.Error(msg)
	} else if l.logLevel == LogWarn {
		l.Warn(msg)
	}
}

func (l *Logger) setOutputToStdout()  {
	l.SetOutput(os.Stdout)
}

func (l *Logger) Info(msg string)  {
	l.Print("[INFO] " + msg)
}

func (l *Logger) Warn(msg string)  {
	l.errorLog("[WARN] " + msg)
}

func (l *Logger) errorLog(msg string) {
	var bf bytes.Buffer
	bf.WriteString(msg)

	_, err := l.errorWriter.Write(bf.Bytes())

	if err != nil {
		l.Print(msg)
	}
}

func (l *Logger) Error(msg string)  {
	l.errorLog("[ERROR] " + msg)
}

func (l *Logger) Debug(msg string) {
	if !l.debug {
		return
	}

	l.errorLog("[DEBUG] " + msg)
}


func NewLogger(name string, logLevel int) *Logger {
	logger := &Logger{
		debug:   false,
		Logger:  log.Logger{},
		name: name,
		logLevel: logLevel,
		errorWriter: os.Stderr,
	}

	logger.SetPrefix(name + " ")
	logger.setOutputToStdout()
	logger.SetFlags(log.LstdFlags)
	return logger
}

