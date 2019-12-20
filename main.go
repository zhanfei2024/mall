package main

import (
	_ "mall/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 参数1   driverName
	// 参数2   数据库类型
	// 这个用来设置 driverName 对应的数据库类型
	// mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 设置默认数据库
	// ORM 必须注册一个别名为 default 的数据库，作为默认使用。
	// 参数1		数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2		driverName
	// 参数3		对应的链接字符串
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))

	// 自动建表
	// 参数1 数据库别名
	// 参数2 true 删除表并重新创建
	// 参数3 打印日志
	err := orm.RunSyncdb("default", false, true)

	if err != nil {
		fmt.Println(err)
	}

}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// 开启日志
	logs.SetLogger(logs.AdapterFile, `{"filename":"test.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)

	// 开启 orm 调试模式：开发过程中建议打开，release时需要关闭
	orm.Debug = true
	beego.Run()
}
