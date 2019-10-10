package log

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	once sync.Once
)

// Init zap log.
func initLog() {
	logger = InitLogger(fmt.Sprintf("%s/%s", viper.GetString("log.path"), "soter-demo.log"), viper.GetString("log.level"))
}

// Get zap log instance.
func Logger() *zap.Logger {
	once.Do(func() {
		initLog()
	})
	return logger
}
