package logs

import (
	"fmt"
	"github.com/cro4k/common/timeutil"
	"github.com/cro4k/gms/layout/public/global"
	"github.com/sirupsen/logrus"
	"strings"
)

type textFormatter struct{}

func (f *textFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var id string
	if entry.Context != nil {
		id = entry.Context.Value("rid").(string)
	}
	var text string
	text += fmt.Sprintf("[%s][%s]", strings.ToUpper(entry.Level.String()[:1]), entry.Time.Format(timeutil.Layout))
	if id != "" {
		text += fmt.Sprintf("[%s]", id)
	}
	if entry.HasCaller() {
		text += fmt.Sprintf(" %s:%d", entry.Caller.File, entry.Caller.Line)
	}
	text += " " + entry.Message
	return []byte(text + "\n"), nil
}

func init() {
	logrus.SetFormatter(&textFormatter{})
	switch global.C().Env {
	case global.Debug:
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetReportCaller(true)
	case global.Develop:
		logrus.SetLevel(logrus.TraceLevel)
		logrus.SetReportCaller(true)
	case global.Produce:
		logrus.SetLevel(logrus.ErrorLevel)
	}
}
