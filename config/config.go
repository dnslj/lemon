package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"lemon/utils/logging"
	"strings"
)

type Config struct {
	Name   string
	//Logger *logrus.Logger
}

// 初始化
func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	logging.Init()
	c.watchConfig()

	return nil
}

// 初始化配置
func (c *Config) initConfig() error {
	if c.Name != "" {
		// 如果指定了配置文件，则解析指定的配置文件
		viper.SetConfigFile(c.Name)
	} else {
		// 如果没有指定配置文件，则解析默认的配置文件
		viper.AddConfigPath("config/conf")
		viper.SetConfigName("config")
	}
	// 设置配置文件格式为YAML
	viper.SetConfigType("yaml")
	// 读取匹配的环境变量
	viper.AutomaticEnv()
	// 读取环境变量的前缀为lemon
	viper.SetEnvPrefix("lemon")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		//log.Infof("Config file changed: %s", in.Name)
		logrus.Infof("Config file changed: %s", in.Name)
	})
}
