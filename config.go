package logrusconf

import (
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sillyhatxu/convenient-utils/tcpclient"
	"github.com/sillyhatxu/logrus-client/fieldhook"
	"github.com/sillyhatxu/logrus-client/logstashhook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type Conf struct {
	Level        logrus.Level
	ReportCaller bool
	Field        map[string]string
	LogFormatter *logrus.JSONFormatter
	FileConf     *FileConf
	LogstashConf *LogstashConf
}

type FileConf struct {
	LogFormatter     *logrus.JSONFormatter
	FilePath         string
	WithMaxAge       time.Duration
	WithRotationTime time.Duration
}

type LogstashConf struct {
	LogFormatter *logrus.JSONFormatter
	Address      string
}

func NewDefaultConf() *Conf {
	return &Conf{}
}

func (conf Conf) Initial() {
	fmt.Println("InitialLogConfig :", fmt.Sprintf("%#v", conf))
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(conf.Level)
	logrus.SetReportCaller(conf.ReportCaller)
	logrus.SetFormatter(conf.LogFormatter)
	if len(conf.Field) > 0 {
		logrus.AddHook(&fieldhook.DefaultFieldHook{Field: conf.Field})
	}
	if conf.LogstashConf != nil {
		if conf.LogstashConf.LogFormatter == nil {
			panic(fmt.Sprintf("you must configure LogFormatter in [Conf.LogstashConf]; %#v", conf))
		}
		conn, err := tcpclient.Dial("tcp", conf.LogstashConf.Address)
		if err != nil {
			panic(fmt.Sprintf("net.Dial(tcp, %s); Error : %v", conf.LogstashConf.Address, err))
		}
		hook := logstashhook.New(conn, conf.LogstashConf.LogFormatter)
		logrus.AddHook(hook)
	}
	if conf.FileConf != nil {
		if conf.FileConf.LogFormatter == nil {
			panic(fmt.Sprintf("you must configure LogFormatter in [Conf.LogstashConf]; %#v", conf))
		}
		if conf.FileConf.FilePath == "" {
			panic(fmt.Sprintf("you must configure FilePath in [Conf.LogstashConf]; %#v", conf))
		}
		if !exists(conf.FileConf.FilePath) {
			err := createFolder(conf.FileConf.FilePath)
			if err != nil {
				panic(fmt.Sprintf("create folder error; Error : %v", err))
			}
		}
		infoWriter, err := rotatelogs.New(
			conf.FileConf.FilePath+"info.log.%Y%m%d",
			rotatelogs.WithLinkName(conf.FileConf.FilePath+"info.log"),
			//rotatelogs.WithLinkName(lc.filePath+lc.module+"-info.log"),
			rotatelogs.WithMaxAge(conf.FileConf.WithMaxAge),
			rotatelogs.WithRotationTime(conf.FileConf.WithRotationTime),
		)
		if err != nil {
			panic(fmt.Sprintf("rotatelogs.New [info writer] error; Error : %v", err))
		}
		errorWriter, err := rotatelogs.New(
			conf.FileConf.FilePath+"error.log.%Y%m%d",
			rotatelogs.WithLinkName(conf.FileConf.FilePath+"error.log"),
			rotatelogs.WithMaxAge(conf.FileConf.WithMaxAge),
			rotatelogs.WithRotationTime(conf.FileConf.WithRotationTime),
		)
		if err != nil {
			panic(fmt.Sprintf("rotatelogs.New [error writer] error; Error : %v", err))
		}
		logrus.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				logrus.InfoLevel:  infoWriter,
				logrus.WarnLevel:  infoWriter,
				logrus.ErrorLevel: infoWriter,
			},
			conf.FileConf.LogFormatter,
		))
		logrus.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				logrus.WarnLevel:  errorWriter,
				logrus.ErrorLevel: errorWriter,
			},
			conf.FileConf.LogFormatter,
		))

	}
}

func createFolder(path string) error {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
