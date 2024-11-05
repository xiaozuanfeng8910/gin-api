package config

import (
	"fmt"
	"gin-api/pkg/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string                    `yaml:"server_port" mapstructure:"server_port"`
	Database   map[string]DatabaseConfig `yaml:"database"`
	Redis      RedisConfig               `yaml:"redis"`
	Log        LogConfig                 `yaml:"log"`
}

func InitializeConfig() (*Config, error) {
	env := utils.GetEnv("APP_ENV", "dev")

	// 设置配置文件的前缀和路径
	v := viper.New()
	v.SetConfigName(env)      // 配置文件名(不带扩展名)
	v.AddConfigPath("config") // 配置文件所在目录
	v.SetConfigType("yaml")   // 配置文件类型

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file:%v", err)
	}

	var config Config

	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := v.Unmarshal(&config); err != nil {
			fmt.Printf("重载配置文件失败:%v\n", err)
		}
	})

	// 将配置文件内容解析到Config结构体中
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config:%v", err)
	}

	return &config, nil
}
