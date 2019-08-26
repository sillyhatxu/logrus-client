package filehook

import (
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

type FileConf struct {
	LogFormatter     logrus.Formatter
	FilePath         string
	WithMaxAge       time.Duration
	WithRotationTime time.Duration
}

func (fc FileConf) CreateFileHook(fileName string, writerLevels []logrus.Level) (*lfshook.LfsHook, error) {
	hookWrite, err := rotatelogs.New(
		fc.FilePath+fileName+".log.%Y%m%d",
		rotatelogs.WithLinkName(fc.FilePath+fileName+".log"),
		//rotatelogs.WithLinkName(lc.filePath+lc.module+"-info.log"),
		rotatelogs.WithMaxAge(fc.WithMaxAge),
		rotatelogs.WithRotationTime(fc.WithRotationTime),
	)
	if err != nil {
		return nil, err
	}
	if writerLevels == nil || len(writerLevels) == 0 {
		return nil, fmt.Errorf("writer levels can not be empty")
	}
	writerMap := make(lfshook.WriterMap)
	for _, writerLevel := range writerLevels {
		writerMap[writerLevel] = hookWrite
	}
	return lfshook.NewHook(writerMap, fc.LogFormatter), nil
}
