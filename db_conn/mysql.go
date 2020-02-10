// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-10
// Time: 20:23

package db_conn

import (
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DbClient *xorm.Engine

func InitMysqlClient() (err error) {
	driverName := "mysql"
	DsName := "root:root@tcp(127.0.0.1:3306)/chat?charset=utf8"
	err = errors.New("")
	DbClient, err = xorm.NewEngine(driverName, DsName)
	if nil != err && "" != err.Error() {
		log.Fatal(err.Error())
	}
	DbClient.ShowSQL(true)
	DbClient.SetMaxOpenConns(2)
	//todo
	//err = DbClient.Sync2(new(model.User), new(model.Contact), new(model.Community))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("init data base ok")
	return err
}
