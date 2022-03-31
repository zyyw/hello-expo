package pkg

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/zyyw/hello-expo/config"
	"github.com/zyyw/hello-expo/log"
)

var (
	// Conf variable
	Conf *config.Config
)

func InitConfig(cfg string) error {
	if err := initViper(cfg); err != nil {
		return err
	}

	return nil
}

func initViper(cfg string) error {
	if cfg != "" {
		viper.SetConfigFile(cfg) // if conf.yaml is given, then resolve this given config file
	} else {
		viper.AddConfigPath("config") // set default config file path
		viper.SetConfigName("conf")   // set default config file name
	}
	viper.SetConfigType("yaml") // set default config file name extension as yaml
	viper.AutomaticEnv()        // read env variable

	if err := getNewConfig(); err != nil {
		return err
	}

	watchConfig()

	return nil
}

func getNewConfig() error {
	var err error
	if err = viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	env := viper.GetString(config.Env)

	var subViper *viper.Viper
	if env == "dev" {
		subViper = viper.Sub("dev")
	} else if env == "staging" {
		subViper = viper.Sub("staging")
	} else if env == "pro" {
		subViper = viper.Sub("pro")
	} else {
		return fmt.Errorf("found invalid env variable: %s", env)
	}
	if Conf, err = NewConfig(subViper); err != nil {
		return err
	}

	return nil
}

// monitor config file and do hot loading
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := getNewConfig(); err != nil {
			log.Logger().Sugar().Errorf("getConfigFile error out, err=%v", err)
		}
		log.Logger().Sugar().Infof(fmt.Sprintf("Config file changed: %s", e.Name))
	})
}

// NewConfig Get config file to struct
func NewConfig(cfg *viper.Viper) (*config.Config, error) {
	c := &config.Config{}

	if err := cfg.Unmarshal(c); err != nil {
		return nil, err
	}
	return c, nil
}
