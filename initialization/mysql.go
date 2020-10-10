package initialization

import (
	"github.com/novliang/helm/global"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Mysql() {
	conf := global.HELM_CONF.Mysql
	if db, err := gorm.Open("mysql", conf.Username+":"+conf.Password+"@("+conf.Path+")/"+conf.Dbname+"?"+conf.Config); err != nil {
		panic("MySQL启动异常: " + err.Error())
	} else {
		global.HELM_DB = db
		global.HELM_DB.DB().SetMaxIdleConns(conf.MaxIdleConns)
		global.HELM_DB.DB().SetMaxOpenConns(conf.MaxOpenConns)
		global.HELM_DB.LogMode(conf.LogMode)
	}
}
