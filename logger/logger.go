// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-10
// Time: 20:20

package logger

import (
	"io"
	"os"

	"ginchat/common"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
)

func LogInit(logPath string) {
	log.SetLevel(log.InfoLevel)
	var multiWriter io.Writer
	if common.LogStdout {
		multiWriter = io.MultiWriter(os.Stdout,
			&lumberjack.Logger{
				Filename:   logPath + "newPlatform.log",
				MaxSize:    100,
				MaxBackups: 1,
				MaxAge:     7,
			})
	} else {
		multiWriter = io.MultiWriter(
			&lumberjack.Logger{
				Filename:   logPath + "newPlatform.log",
				MaxSize:    100,
				MaxBackups: 1,
				MaxAge:     7,
			})
	}
	log.SetOutput(multiWriter)
	gin.DefaultWriter = multiWriter
}
