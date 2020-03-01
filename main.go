// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-10
// Time: 20:10

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"ginchat/common"
	"ginchat/db_conn"
	"ginchat/httphandlerpack/contact"
	"ginchat/logger"
	"ginchat/router"
)

func main() {
	// 初始化系统日志
	logger.LogInit(common.LogPath)

	//初始化mysql的客户端
	if err := db_conn.InitMysqlClient(); err != nil {
		fmt.Printf("InitMysqlClient：%s", err.Error())
		os.Exit(1)
	}

	//初始化Redis的客户端
	if err := db_conn.InitRedisClient(); err != nil {
		fmt.Printf("InitRedisClient：%s", err.Error())
		os.Exit(1)
	}

	//群组消息初始化
	if err := contact.InitCommunityRedis(); err != nil {
		fmt.Printf("InitCommunityRedis：%s", err.Error())
		os.Exit(1)
	}

	if len(common.Nodes) > 1 {
		go router.ListenRpc()
	}
	r, _ := router.Register()
	r.Use(gin.Logger())
	_ = r.Run(":" + strconv.Itoa(common.ServerPort))
}
