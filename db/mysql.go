package db


import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/TriangleGo/gorouter/config"
)

var _flag string

func InitMysql() {

	
	dbhost := config.GetConfig("mysql_host")
	dbuser := config.GetConfig("mysql_user")
	dbpass := config.GetConfig("mysql_pass")
	dbname := config.GetConfig("mysql_dbname")
	
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//username:password@protocol(address)/dbname?param=value
	conn := dbuser + ":" + dbpass + "@tcp(" + dbhost + ")/" + dbname + "?charset=utf8"
	//fmt.Printf("ConnString %v\n",conn)
	orm.RegisterDataBase("default", "mysql", conn)
	
}


func GetDB() orm.Ormer{
	if _flag == "" {
		InitMysql()
		_flag = "ready"
	}
	return orm.NewOrm()
}