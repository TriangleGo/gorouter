package db

import (
	"time"
	"fmt"
	"testing"
	_"github.com/astaxie/beego/orm"
)



func Test_DB(t *testing.T) {
	InitMysql()
	
	for {
		var Enterpriseno int
		err := GetDB().Raw("select enterpriseno from im_config where configid = ?",16).QueryRow(&Enterpriseno)
	
		fmt.Printf("sql query err = %v \nvalue=%v\n",err,Enterpriseno)
		time.Sleep( 1 * time.Second)
	}

}

