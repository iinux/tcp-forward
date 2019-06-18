package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iinux/tcp-forward/conf"
	"github.com/iinux/tcp-forward/statistic"
	"strings"
)

func HttpServer()  {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "statistic",
			"data": statistic.Get(),
		})
	})

	// 路由组使用 gin.BasicAuth() 中间件
	// gin.Accounts 是 map[string]string 的一种快捷方式
	authorized := router.Group("/admin", gin.BasicAuth(conf.Account))

	// /admin/secrets 端点
	// 触发 "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		remoteAddr := c.Request.RemoteAddr
		remoteHost := strings.Split(remoteAddr, ":")[0]
		conf.Container.GetWhitelist().AddIp(remoteHost)
		c.JSON(http.StatusOK, gin.H{"code":0,"user":user})
	})

	router.Run(conf.HttpPort)
}
