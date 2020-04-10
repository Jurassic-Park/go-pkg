package logrus

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logging *logrus.Logger

// Setup initializes the logrus instance
func Setup() {
	Logging = logrus.New()

	// 为当前logrus实例设置消息的输出，同样地，
	// 可以设置logrus实例的输出到任意io.writer
	Logging.Out = os.Stdout

	// 为当前logrus实例设置消息输出格式为json格式。
	// 同样地，也可以单独为某个logrus实例设置日志级别和hook，这里不详细叙述。
	Logging.Formatter = &logrus.JSONFormatter{}
}
