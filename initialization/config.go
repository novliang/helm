package initialization

import (
	"fmt"
	"github.com/novliang/helm/global"
	"github.com/spf13/viper"
)

const defaultConfigFile = "config.yaml"

func Config() {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if err := v.Unmarshal(&global.HELM_CONF); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	global.HELM_VIPER = v
}
