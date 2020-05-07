package logger

import (
	"io"
	"os"
)

type FileLogger struct {
	logFile io.Writer
	errorLogFIle io.Writer
	Logger
}

func (fl *FileLogger) Info(msg string) {
	fl.SetOutput(fl.logFile)
	fl.Logger.Info(msg)
}

func (fl *FileLogger) Error(msg string) {
	fl.SetOutput(fl.errorLogFIle)
	fl.Logger.Error(msg)
}

func (fl *FileLogger) Debug(msg string) {
	fl.SetOutput(fl.errorLogFIle)
	fl.Logger.Debug(msg)
}

func (fl *FileLogger) Warn(msg string) {
	fl.SetOutput(fl.errorLogFIle)
	fl.Logger.Warn(msg)
}

func NewFileLogger(name string, outputFile string, logLevel int) *FileLogger  {
	infoLogFile, err := os.OpenFile(outputFile + "_info", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0644)
	errorLogFile, errorLogOpenErr := os.OpenFile(outputFile + "_error", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	if errorLogOpenErr != nil {
		panic(errorLogOpenErr)
	}

	fileLogger := &FileLogger{
		logFile: infoLogFile,
		errorLogFIle: errorLogFile,
		Logger:  Logger{name: name, logLevel: logLevel, errorWriter: errorLogFile},
	}

	fileLogger.SetOutput(infoLogFile)
	return fileLogger
}
