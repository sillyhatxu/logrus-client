package fieldhook

import "github.com/sirupsen/logrus"

type DefaultFieldHook struct {
	Field map[string]string
}

func (h *DefaultFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *DefaultFieldHook) Fire(e *logrus.Entry) error {
	for k, v := range h.Field {
		e.Data[k] = v
	}
	return nil
}
