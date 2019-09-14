package logrusconf

import (
	"github.com/sillyhatxu/logrus-client/filehook"
	"github.com/sillyhatxu/logrus-client/logstashhook"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

var defaultJSONFormatter = &logrus.JSONFormatter{
	TimestampFormat: time.RFC3339Nano,
	FieldMap: *&logrus.FieldMap{
		logrus.FieldKeyMsg:  "message",
		logrus.FieldKeyTime: "@timestamp",
	},
}

var defaultFields = logrus.Fields{
	"project":  "test",
	"module":   "test-module",
	"@version": "1",
	"type":     "project-log",
}

func TestDefaultLog(t *testing.T) {
	var config = NewLogrusConfig(
		Fields(defaultFields),
	)
	config.Initial()
	logrus.Debugf("test debug log[%s]", "This is debug log")
	logrus.Infof("test info log[%s]", "This is info log")
	logrus.Errorf("test error log[%s]", "This is error log")
	logrus.Warningf("test warn log[%s]", "This is warn log")
}

func TestFileLog(t *testing.T) {
	var fileConfig = filehook.NewFileConfig("/Users/shikuanxu/go/src/github.com/sillyhatxu/logrus-client/logs/")
	var config = NewLogrusConfig(
		Fields(defaultFields),
		FileConfig(fileConfig),
	)
	config.Initial()
	logrus.Debugf("test debug log[%s]", "This is debug log")
	logrus.Infof("test info log[%s]", "This is info log")
	logrus.Errorf("test error log[%s]", "This is error log")
	logrus.Warningf("test warn log[%s]", "This is warn log")
}

func TestLogstashLog(t *testing.T) {
	var logstashConfig = logstashhook.NewLogstashConfig("localhost:5000", logstashhook.Fields(defaultFields))
	var config = NewLogrusConfig(
		Fields(defaultFields),
		LogstashConfig(logstashConfig),
	)
	config.Initial()
	//logrus.Debugf("test debug log[%s]", "This is debug log")
	logrus.Infof("test info log[%s]", "This is info log")
	//logrus.Errorf("test error log[%s]", "This is error log")
	//logrus.Warningf("test warn log[%s]", "This is warn log")
}

func TestInputLogstash(t *testing.T) {
	var logstashConfig = logstashhook.NewLogstashConfig("localhost:5000", logstashhook.Fields(defaultFields))
	var config = NewLogrusConfig(
		Fields(defaultFields),
		LogstashConfig(logstashConfig),
	)
	config.Initial()
	var i int64
	for {
		logrus.Infof("test info[%d] %v", i, time.Now().Format("2006-01-02 15:04:05"))
		i++
		time.Sleep(5 * time.Second)
	}
}
