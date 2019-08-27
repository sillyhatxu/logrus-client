package logrusconf

import (
	"github.com/sillyhatxu/logrus-client/filehook"
	"github.com/sillyhatxu/logrus-client/logstashhook"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestDefaultLog(t *testing.T) {
	jsonFormatter := &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		FieldMap: *&logrus.FieldMap{
			logrus.FieldKeyMsg:  "message",
			logrus.FieldKeyTime: "@timestamp",
		},
	}
	conf := &Conf{
		Level:        logrus.InfoLevel,
		ReportCaller: true,
		Fields: logrus.Fields{
			"project": "test",
			"module":  "test-module",
		},
		LogFormatter: jsonFormatter,
	}
	conf.Initial()
	logrus.Debugf("test debug log[%s]", "This is debug log")
	logrus.Infof("test info log[%s]", "This is info log")
	logrus.Errorf("test error log[%s]", "This is error log")
	logrus.Warningf("test warn log[%s]", "This is warn log")
}

func TestFileLog(t *testing.T) {
	jsonFormatter := &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		FieldMap: *&logrus.FieldMap{
			logrus.FieldKeyMsg:  "message",
			logrus.FieldKeyTime: "@timestamp",
		},
	}
	conf := &Conf{
		Level:        logrus.InfoLevel,
		ReportCaller: true,
		Fields: logrus.Fields{
			"project": "test",
			"module":  "test-module",
		},
		LogFormatter: jsonFormatter,
		FileConf: &filehook.FileConf{
			LogFormatter: &logrus.TextFormatter{
				DisableColors:    true,
				FullTimestamp:    true,
				TimestampFormat:  string("2006-01-02 15:04:05"),
				QuoteEmptyFields: true,
				FieldMap: *&logrus.FieldMap{
					logrus.FieldKeyMsg:  "message",
					logrus.FieldKeyTime: "timestamp",
				},
			},
			FilePath:         "/Users/cookie/go/gopath/src/github.com/sillyhatxu/logrus-client/logs/",
			WithMaxAge:       time.Duration(876000) * time.Hour,
			WithRotationTime: time.Duration(24) * time.Hour,
		},
	}
	conf.Initial()
	logrus.Debugf("test debug log[%s]", "This is debug log")
	logrus.Infof("test info log[%s]", "This is info log")
	logrus.Errorf("test error log[%s]", "This is error log")
	logrus.Warningf("test warn log[%s]", "This is warn log")
}

func TestLogstashLog(t *testing.T) {
	jsonFormatter := &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		FieldMap: *&logrus.FieldMap{
			logrus.FieldKeyMsg:  "message",
			logrus.FieldKeyTime: "@timestamp",
		},
	}
	conf := &Conf{
		Level:        logrus.InfoLevel,
		ReportCaller: true,
		Fields: logrus.Fields{
			"project": "test",
			"module":  "test-module",
		},
		LogFormatter: jsonFormatter,
		LogstashConf: &logstashhook.LogstashConf{
			LogFormatter: jsonFormatter,
			Address:      "localhost:5000",
			ExtraFields:  logrus.Fields{"@version": "1", "type": "project-log"},
		},
	}
	conf.Initial()
	logrus.Debugf("test debug log[%s]", "This is debug log")
	logrus.Infof("test info log[%s]", "This is info log")
	logrus.Errorf("test error log[%s]", "This is error log")
	logrus.Warningf("test warn log[%s]", "This is warn log")
}

func TestInputLogstash(t *testing.T) {
	jsonFormatter := &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		FieldMap: *&logrus.FieldMap{
			logrus.FieldKeyMsg:  "message",
			logrus.FieldKeyTime: "@timestamp",
		},
	}
	conf := &Conf{
		Level:        logrus.InfoLevel,
		ReportCaller: true,
		Fields: logrus.Fields{
			"project": "test",
			"module":  "test-module",
		},
		LogFormatter: jsonFormatter,
		FileConf: &filehook.FileConf{
			LogFormatter: &logrus.TextFormatter{
				DisableColors:    true,
				FullTimestamp:    true,
				TimestampFormat:  string("2006-01-02 15:04:05"),
				QuoteEmptyFields: true,
				FieldMap: *&logrus.FieldMap{
					logrus.FieldKeyMsg:  "message",
					logrus.FieldKeyTime: "timestamp",
				},
			},
			FilePath:         "/Users/cookie/go/gopath/src/github.com/sillyhatxu/logrus-client/logs",
			WithMaxAge:       time.Duration(876000) * time.Hour,
			WithRotationTime: time.Duration(24) * time.Hour,
		},
		LogstashConf: &logstashhook.LogstashConf{
			LogFormatter: jsonFormatter,
			Address:      "localhost:51401",
		},
	}
	conf.Initial()
	var i int64
	for {
		logrus.Infof("test info[%d] %v", i, time.Now().Format("2006-01-02 15:04:05"))
		i++
		time.Sleep(5 * time.Second)
	}
}
