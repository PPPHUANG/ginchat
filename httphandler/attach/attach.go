// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-11
// Time: 01:21

package attach

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"ginchat/common"
	"ginchat/util"

	log "github.com/sirupsen/logrus"
)

func init() {
	err := os.MkdirAll(common.AttachPath, os.ModePerm)
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/httphandler/attach.go",
		}).Error("Make Attach Dir Error")
	}
}
func Upload(c *gin.Context) {
	UploadLocal(c)
	//UploadOss(w,r)
}

//url格式 /mnt/xxxx.png  需要确保网络能访问/mnt/
func UploadLocal(c *gin.Context) {
	srcFile, err := c.FormFile("file")
	if err != nil {
		log.Error("Upload File Error")
		util.RespFail(c, err.Error())
		return
	}
	suffix := ".png"
	//如果前端文件名称包含后缀 xx.xx.png
	oFilename := srcFile.Filename
	tmp := strings.Split(oFilename, ".")
	if len(tmp) > 1 {
		suffix = "." + tmp[len(tmp)-1]
	}
	//如果前端指定filetype
	//formdata.append("filetype",".png")
	fileType := c.PostForm("filetype")
	if len(fileType) > 0 {
		suffix = fileType
	}
	filename := fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
	if err != nil {
		util.RespFail(c, err.Error())
		return
	}
	//将源文件内容存储到新的文件夹
	err = c.SaveUploadedFile(srcFile, "./mnt/"+filename)
	if err != nil {
		log.WithFields(log.Fields{
			"filename": "/attach.go/UploadLocal",
		}).Error(err.Error())
		util.RespFail(c, err.Error())
		return
	}
	//将新文件路径转换成url地址
	url := "/mnt/" + filename
	util.RespOk(c, url, "")
}
