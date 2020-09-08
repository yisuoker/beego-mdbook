package sysinit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func dbinit(aliases ...string) {
	// 开发模式，显示命令信息
	isDev := beego.AppConfig.String("runmode") == "dev"
	if isDev {
		orm.Debug = isDev
	}

	// 注册数据库
	if len(aliases) > 0 {
		for _, alias := range aliases {
			registerDatabase(alias)
		}
	} else {
		registerDatabase("w")
	}
}

// alias 数据库别名
func registerDatabase(alias string) {
	if len(alias) == 0 {
		return
	}

	// 连接别名
	dbAlias := alias
	if "w" == alias {
		dbAlias = "default" // 默认
	}

	dbHost := beego.AppConfig.String("db_w_host")
	dbPort := beego.AppConfig.String("db_w_port")
	dbUser := beego.AppConfig.String("db_w_user")
	dbPwd := beego.AppConfig.String("db_w_password")
	dbName := beego.AppConfig.String("db_w_database")
	dbCharset := beego.AppConfig.String("db_w_charset")
	ds := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=" + dbCharset

	orm.RegisterDataBase(dbAlias, "mysql", ds, 10, 10)

	// 自动建表
	if "w" == alias && "default" == dbAlias {
		orm.RunSyncdb(dbAlias, false, true)
	}
}
