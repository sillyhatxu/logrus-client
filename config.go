package logrusconf

import (
	"fmt"
	"github.com/sillyhatxu/logrus-client/fieldhook"
	"github.com/sillyhatxu/logrus-client/filehook"
	"github.com/sillyhatxu/logrus-client/logstashhook"
	"github.com/sirupsen/logrus"
	"os"
)

type Conf struct {
	Level        logrus.Level
	ReportCaller bool
	Fields       logrus.Fields
	LogFormatter logrus.Formatter
	FileConf     *filehook.FileConf
	LogstashConf *logstashhook.LogstashConf
}

func (conf Conf) validate() error {
	if conf.FileConf != nil {
		//if conf.FileConf.LogFormatter == nil {
		//	return fmt.Errorf("you must configure LogFormatter in [Conf.LogstashConf]; %#v", conf)
		//}
		if conf.FileConf.FilePath == "" {
			return fmt.Errorf("you must configure FilePath in [Conf.LogstashConf]; %#v", conf)
		}
		if !exists(conf.FileConf.FilePath) {
			err := createFolder(conf.FileConf.FilePath)
			if err != nil {
				return fmt.Errorf("create folder error; Error : %v", err)
			}
		}
	}
	if conf.LogstashConf != nil {
		//if conf.LogstashConf.LogFormatter == "" {
		//	return fmt.Errorf("you must configure LogFormatter in [Conf.LogstashConf]; %#v", conf)
		//}
		if conf.LogstashConf.Address == "" {
			return fmt.Errorf("you must configure address in [Conf.LogstashConf]; %#v", conf)
		}
	}
	return nil
}

func (conf Conf) Initial() {
	fmt.Println("InitialLogConfig :", fmt.Sprintf("%#v", conf))
	err := conf.validate()
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(conf.Level)
	logrus.SetReportCaller(conf.ReportCaller)
	logrus.SetFormatter(conf.LogFormatter)
	if len(conf.Fields) > 0 {
		logrus.AddHook(&fieldhook.DefaultFieldHook{Fields: conf.Fields})
	}
	if conf.FileConf != nil {
		infoHook, err := conf.FileConf.CreateFileHook("info", []logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel})
		if err != nil {
			panic(fmt.Sprintf("create info hook error; Error : %v", err))
		}
		logrus.AddHook(infoHook)
		errorHook, err := conf.FileConf.CreateFileHook("error", []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel})
		if err != nil {
			panic(fmt.Sprintf("create error hook error; Error : %v", err))
		}
		logrus.AddHook(errorHook)
	}
	if conf.LogstashConf != nil {
		hook := conf.LogstashConf.New(conf.Fields)
		logrus.AddHook(hook)
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
