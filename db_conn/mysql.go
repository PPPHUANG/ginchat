// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-10
// Time: 20:23

package db_conn

import (
	"fmt"

	"ginchat/common"
	"ginchat/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	log "github.com/sirupsen/logrus"
)

var DbClient *xorm.Engine

func InitMysqlClient() (err error) {
	driverName := "mysql"
	DsName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", common.MysqlUser, common.MysqlPwd, common.MysqlIp, common.MysqlPort, common.DBName)
	DbClient, err = xorm.NewEngine(driverName, DsName)
	if err != nil {
		panic(err.Error())
	}
	DbClient.ShowSQL(common.ShowSQL)
	DbClient.SetMaxOpenConns(common.MaxOpenConns)
	err = DbClient.Sync2(new(model.User), new(model.Contact), new(model.Community))
	if err != nil {
		panic(err.Error())
	}
	log.WithFields(log.Fields{
		"filename": "/common/mysql.go",
	}).Info("Init Database Successfully")
	return err
}
