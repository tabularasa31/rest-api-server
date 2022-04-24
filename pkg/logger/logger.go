package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error { // формирует запись в лог
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, m := range hook.Writer {
		m.Write([]byte(line))
	}
	return nil
}

func (hook *writerHook) Levels() []logrus.Level { // возвращает levels из хука
	return hook.LogLevels
}

// ---------------
var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func (l *Logger) GetLoggerWithFields(k string, v interface{}) Logger { // если есть необходимость создать другой логгер с еще полем
	return Logger{l.WithField(k, v)}
}

// ---------------

func init() {
	l := logrus.New()
	l.SetReportCaller(true) //sets whether the standard logger will include the calling method as a field
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line) // в какой функции происходит логирование и в какой строке
		},
		DisableColors: false, // Force disabling colors
		FullTimestamp: true,  // Disable timestamp logging. useful when output is redirected to logging system that already adds timestamps.
	}
	err := os.MkdirAll("logs", 0755)
	if err != nil {
		panic(err)
	}

	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	l.SetOutput(io.Discard) // all Write calls succeed without doing anything - чтобы по умолчанию логи никуда не уходили

	l.AddHook(&writerHook{
		Writer:    []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}
