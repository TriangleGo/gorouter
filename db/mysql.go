package db


import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var _flag string

func InitMysql() {
	var	dbhost string
	var	dbuser string
	var	dbpass string
	var	dbname string
	
	dbhost="127.0.0.1:51816"
	dbuser="oa_local"
	dbpass=`f*(&Dssdsa)s`
	dbname="db_oa_enterprise"
	
	orm.RegisterDriver("mysql", orm.DR_MySQL)
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