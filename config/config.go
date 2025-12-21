package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote" 
)

var (
	Service      service
	runtimeViper = viper.New()
)

func Init() {
	port := os.Getenv("PORT")
	if port == "" {
		logger.Fatalf("config.Init: service port is empty")
	}
	Service.Addr = fmt.Sprintf("0.0.0.0:%s", port)

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		logger.Fatalf("config.Init: config path is empty")
	}

	runtimeViper.SetConfigFile(configPath)
	if err := runtimeViper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			logger.Fatal("config.Init: could not find config files")
		}
		logger.Fatalf("config.Init: read config error: %v", err)
	}

	configMapping()

	runtimeViper.OnConfigChange(func(in fsnotify.Event) {
		logger.Infof("config: notice config changed: %v\n", in.String())
		configMapping() // 重新映射配置
	})
	runtimeViper.WatchConfig()
}

func configMapping() {
	c := new(config)
	if err := runtimeViper.Unmarshal(&c); err != nil {
		// 由于这个函数会在配置重载时被再次触发，所以需要判断日志记录方式
		logger.Fatalf("config.configMapping: config: unmarshal error: %v", err)
	}
}
