package log

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
	"github.com/zyyw/hello-expo/config"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	once   sync.Once
)

// Init zap log.
func initLog() {
	logPath := fmt.Sprintf("%s.log.path", viper.GetString(config.Env))
	logLevel := fmt.Sprintf("%s.log.level", viper.GetString(config.Env))
	logger = InitLogger(fmt.Sprintf("%s/%s", viper.GetString(logPath), "hello-expo.log"), viper.GetString(logLevel))
}

// Get zap log instance.
func Logger() *zap.Logger {
	once.Do(func() {
		initLog()
	})
	return logger
}
