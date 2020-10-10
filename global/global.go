package global

import (
	"github.com/novliang/helm/config"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

var (
	HELM_DB    *gorm.DB
	HELM_CONF  *config.Config
	HELM_LOGGER   *log.Logger
	HELM_VIPER *viper.Viper
)
