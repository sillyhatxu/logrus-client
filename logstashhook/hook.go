package logstashhook

import (
	"fmt"
	"github.com/sillyhatxu/logrus-client/tcpclient"
	"github.com/sirupsen/logrus"
	"io"
	"sync"
)

type LogstashConf struct {
	LogFormatter logrus.Formatter
	ExtraFields  logrus.Fields
	Address      string
}

type Hook struct {
	writer    io.Writer
	formatter logrus.Formatter
}

func (lc LogstashConf) New(fields logrus.Fields) logrus.Hook {
	for k, v := range lc.ExtraFields {
		if _, ok := fields[k]; !ok {
			fields[k] = v
		}
	}
	conn, err := tcpclient.Dial("tcp", lc.Address)
	if err != nil {
		panic(fmt.Sprintf("net.Dial(tcp, %s); Error : %v", lc.Address, err))
	}
	return Hook{
		writer: conn,
		formatter: LogstashFormatter{
			Formatter: lc.LogFormatter,
			Fields:    fields,
		},
	}
}

func (h Hook) Fire(e *logrus.Entry) error {
	dataBytes, err := h.formatter.Format(e)
	if err != nil {
		return err
	}
	_, err = h.writer.Write(dataBytes)
	return err
}

func (h Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

//var (
//	logstashFields   = logrus.Fields{"@version": "1", "type": "project-log"}
//	logstashFieldMap = logrus.FieldMap{
//		logrus.FieldKeyTime: "@timestamp",
//		logrus.FieldKeyMsg:  "message",
//	}
//)

//func DefaultFormatter(fields logrus.Fields) logrus.Formatter {
//	for k, v := range logstashFields {
//		if _, ok := fields[k]; !ok {
//			fields[k] = v
//		}
//	}
//
//	return LogstashFormatter{
//		Formatter: &logrus.JSONFormatter{FieldMap: logstashFieldMap},
//		Fields:    fields,
//	}
//}

type LogstashFormatter struct {
	logrus.Formatter
	logrus.Fields
}

func (f LogstashFormatter) Format(e *logrus.Entry) ([]byte, error) {
	ne := copyEntry(e, f.Fields)
	dataBytes, err := f.Formatter.Format(ne)
	releaseEntry(ne)
	return dataBytes, err
}

func copyEntry(e *logrus.Entry, fields logrus.Fields) *logrus.Entry {
	ne := entryPool.Get().(*logrus.Entry)
	ne.Message = e.Message
	ne.Level = e.Level
	ne.Time = e.Time
	ne.Data = logrus.Fields{}
	for k, v := range fields {
		ne.Data[k] = v
	}
	for k, v := range e.Data {
		ne.Data[k] = v
	}
	return ne
}

var entryPool = sync.Pool{
	New: func() interface{} {
		return &logrus.Entry{}
	},
}

func releaseEntry(e *logrus.Entry) {
	entryPool.Put(e)
}
