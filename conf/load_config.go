package conf

import (
	"fmt"

	"github.com/TRON-US/soter-demo/log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig(cfg string) error {
	if err := initViper(cfg); err != nil {
		return err
	}
	return nil
}

func initViper(cfg string) error {
	if cfg != "" {
		viper.SetConfigFile(cfg)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	watchConfig()

	return nil
}

func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Logger().Info(fmt.Sprintf("Config file changed: %s", e.Name))
	})
}
