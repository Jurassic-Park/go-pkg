package go_pkg

import (
	"github.com/Jurassic-Park/go-pkg/gorm"
	"github.com/Jurassic-Park/go-pkg/gredis"
	"github.com/Jurassic-Park/go-pkg/logrus"
	"github.com/Jurassic-Park/go-pkg/setting"
)

// Setup ...
func Setup() {
	setting.Setup()
	logrus.Setup()
	gorm.Setup()
	gredis.Setup()
}
