package initialization

import (
	"github.com/labstack/gommon/log"
	"github.com/novliang/helm/global"
)

func Logger() {
	c := global.HELM_CONF.Log
	logger := log.New(c.Prefix)
	global.HELM_LOGGER = logger
}
