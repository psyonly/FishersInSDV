package settings

import (
	"fmt"
)

import (
	"github.com/spf13/viper"
)

// Config 配置文件结构体
type Config struct {
	AppName string `mapstructure:"appname"`
	Version string `mapstructure:"version"`
	Port    int    `mapstructure:"port"`
}

var (
	// 全局变量Conf从配置文件读取到里面
	Conf = new(Config)
)

func InitSettings() error {
	viper.SetConfigFile("../api/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = viper.Unmarshal(Conf)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Config init success")
	return nil
}
