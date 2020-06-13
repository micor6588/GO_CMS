package sysinit

import (
	_ "GO_CMS/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func initDB() {
	// 数据库连接名称
	dbAlias := beego.AppConfig.String("db_alias")
	// 数据库名称
	dbName := beego.AppConfig.String("db_name")
	// 数据库连接的用户名
	dbUser := beego.AppConfig.String("db_user")
	// 数据库密码
	dbPwd := beego.AppConfig.String("db_pwd")
	// 数据库链接的域名（IP)
	dbHost := beego.AppConfig.String("db_host")
	// 数据库端口：
	dbPort := beego.AppConfig.String("db_port")
	// 数据库的编码格式
	dbCharset := beego.AppConfig.String("db_charset")
	orm.RegisterDataBase(dbAlias, "mysql", dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset="+dbCharset, 30)
	// 如果是开发模式，则显示命令信息
	isDev := (beego.AppConfig.String("runmode") == "dev")
	// 自动建表
	orm.RunSyncdb("default", false, isDev)
	if isDev {
		orm.Debug = isDev
	}
}
