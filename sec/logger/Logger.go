package logger

import (
	"github.com/meowalien/rabbitgather-article/sec/conf"
	"github.com/meowalien/rabbitgather-lib/text"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)


func InitLogger() {
	Logger = CreateLogger()
}

func CreateLogger() *logrus.Entry {
	logger := logrus.New()
	logger.SetFormatter(&MyFormatter{
		ColorEncoding: conf.DEBUG_MOD,
		Formatter: &logrus.JSONFormatter{
			// time格式
			TimestampFormat: time.StampNano,
			PrettyPrint:     true,
		}})
	//}

	logger.SetReportCaller(true)
	//輸出終端機
	logger.SetOutput(io.MultiWriter(os.Stdout))
	//設定log等級
	logger.SetLevel(logrus.DebugLevel)
	return logger.WithFields(logrus.Fields{
		"module": "article",
	})
}

type MyFormatter struct {
	logrus.Formatter
	ColorEncoding bool
}

func (f *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	b, e := f.Formatter.Format(entry)
	if !f.ColorEncoding {
		return b, e
	}
	switch entry.Level {
	case logrus.PanicLevel:
		fallthrough
	case logrus.FatalLevel:
		fallthrough
	case logrus.ErrorLevel:
		return text.ColorByteSting(b, text.FgRed), e
	case logrus.WarnLevel:
		return text.ColorByteSting(b, text.FgYellow), e
	case logrus.InfoLevel:
		return text.ColorByteSting(b, text.FgGreen), e
	case logrus.DebugLevel:
		return text.ColorByteSting(b, text.FgBlue), e
	default:
		return b, e
	}

}

var Logger *logrus.Entry
