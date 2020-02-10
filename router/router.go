// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-10
// Time: 20:32

package router

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"

	"ginchat/common"
	"ginchat/httphandler/attach"
	"ginchat/httphandler/chat"
	"ginchat/httphandler/contact"
	"ginchat/httphandler/user"
	"ginchat/version"
)

func Register() (*gin.Engine, error) {
	if common.ServerDebug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	//指定目录的静态文件
	//http.Handle("/asset/", http.FileServer(http.Dir(".")))
	//http.Handle("/mnt/", http.FileServer(http.Dir(".")))
	router.Static("/assets/", "./assets")
	router.Static("/mnt/", "./mnt")

	RegisterView(router)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/v1/version", version.GetVersion)
	router.POST("/v1/version", version.GetVersion)

	if gin.Mode() == gin.DebugMode {
		router.Use(validateDebug()) // 使用validateDebug()中间件身份验证 以上路由不需要验证
	} else {
		router.Use(validateRelease()) // 使用validateRelease()中间件身份验证 以上路由不需要验证
	}

	userRT := router.Group("/user")
	{
		userRT.POST("/login", user.UserLogin)
		userRT.POST("/register", user.UserRegister)
	}

	contactRT := router.Group("/contact")
	{
		contactRT.POST("/loadcommunity/", contact.LoadCommunity)
		contactRT.POST("/loadfriend", contact.LoadFriend)
		contactRT.GET("/joincommunity", contact.JoinCommunity)
		contactRT.GET("/createcommunity", contact.CreateCommunity)
		contactRT.POST("/addfriend", contact.AddFriend)
	}

	chatRT := router.Group("/chat")
	{
		chatRT.POST("", chat.Chat)
	}
	attachRT := router.Group("/attach")
	{
		attachRT.POST("/upload", attach.Upload)
	}
	return router, nil
}

func validateDebug() gin.HandlerFunc {
	return func(c *gin.Context) {
		if true {
			c.Next() /*该句可以省略，写出来只是表明可以进行验证下一步中间件>，不写，也是内置会继续访问下一个中间件的*/
		} else {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "authorization failed"})
			return // return也是可以省略的，执行了abort操作，会内置在中间件defer前，return，写出来也只是解答为什么Abort()之后，还能执行返回JSON数据
		}
	}
}

func validateRelease() gin.HandlerFunc {
	return func(c *gin.Context) {
		if true {
			c.Next() /*该句可以省略，写出来只是表明可以进行验证下一步中间件>，不写，也是内置会继续访问下一个中间件的*/
		} else {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusBadRequest, "msg": "authorization failed"})
			return // return也是可以省略的，执行了abort操作，会内置在中间件defer前，return，写出来也只是解答为什么Abort()之后，还能执行返回JSON数据
		}
	}
}

func RegisterView(router *gin.Engine) {
	//一次解析出全部模板
	tpl, err := template.ParseGlob("view/**/*")
	if nil != err {
		log.Fatal(err)
	}
	//通过for循环做好映射
	for _, v := range tpl.Templates() {
		tplname := v.Name()
		fmt.Println("Template    " + v.Name())
		router.GET(tplname, func(ctx *gin.Context) {
			fmt.Println("parse  " + v.Name() + "==" + tplname)
			err := tpl.ExecuteTemplate(ctx.Writer, tplname, nil)
			if err != nil {
				log.Fatal(err.Error())
			}
		})
	}

}
