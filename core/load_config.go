package core

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
		viper.SetConfigFile(cfg) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("conf") // 设置配置文件路径
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml") // 设置配置文件格式为YAML
	viper.AutomaticEnv()        // 读取匹配的环境变量

	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}
	watchConfig()

	return nil
}

// 监控配置文件变化并热加载程序
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Logger().Info(fmt.Sprintf("Config file changed: %s", e.Name))
	})
}

