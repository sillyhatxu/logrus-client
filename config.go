package logrusconf

import (
	"fmt"
	"github.com/sillyhatxu/logrus-client/constants"
	"github.com/sillyhatxu/logrus-client/fieldhook"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func NewLogrusConfig(opts ...Option) *Config {
	//default
	config := &Config{
		reportCaller:   true,
		level:          logrus.InfoLevel,
		fields:         nil,
		logFormatter:   constants.DefaultTextFormatter,
		fileConfig:     nil,
		logstashConfig: nil,
	}

	for _, opt := range opts {
		opt(config)
	}
	return config
}

func (c Config) validate() error {
	return nil
}

func (c Config) Initial() {
	log.Println(fmt.Sprintf("InitialLogConfig : %#v", c))
	err := c.validate()
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(c.level)
	logrus.SetReportCaller(c.reportCaller)
	logrus.SetFormatter(c.logFormatter)
	if c.fields != nil && len(c.fields) > 0 {
		logrus.AddHook(&fieldhook.DefaultFieldHook{Fields: c.fields})
	}
	err = c.initialFileConfig()
	if err != nil {
		panic(err)
	}
	err = c.initialLogstashConf()
	if err != nil {
		panic(err)
	}
}

func (c Config) initialFileConfig() error {
	if c.fileConfig == nil {
		return nil
	}
	infoHook, err := c.fileConfig.CreateFileHook("info", []logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel})
	if err != nil {
		return fmt.Errorf("create info hook error; Error : %v", err)
	}
	logrus.AddHook(infoHook)
	errorHook, err := c.fileConfig.CreateFileHook("error", []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel})
	if err != nil {
		return fmt.Errorf("create file hook error; Error : %v", err)
	}
	logrus.AddHook(errorHook)
	return nil
}

func (c Config) initialLogstashConf() error {
	if c.logstashConfig == nil {
		return nil
	}
	hook, err := c.logstashConfig.CreateLogstashHook()
	if err != nil {
		return fmt.Errorf("create logstash hook error; Error : %v", err)
	}
	logrus.AddHook(hook)
	return nil
}
