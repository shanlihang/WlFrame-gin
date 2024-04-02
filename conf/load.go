package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadResource() *Config {
	Conf := new(Config)
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./conf/")

	if readErr := v.ReadInConfig(); readErr != nil {
		panic(fmt.Sprintf("配置文件读取出错,%s\n", readErr))
	}
	if bindErr := v.Unmarshal(Conf); bindErr != nil {
		panic(fmt.Sprintf("配置绑定出错,%s\n", bindErr))
	}

	return Conf
}
