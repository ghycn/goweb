package main

import (
	"github.com/gin-gonic/gin"
	db "goweb/database"
	"goweb/models"
	"log"
	"net/http"
	"time"
)

func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":9090")
}

/**
 * 主程序入口
 */
//func main() {
//	router := gin.Default()
//	router.LoadHTMLGlob("com/view/*")
//
//	router.Use(StatCost())
//	router.GET("/", json)
//	router.GET("/index", html)
//	err := router.Run(":9090")
//	if err != nil {
//		fmt.Println("服务器启动失败！")
//	}
//}

func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		//可以设置一些公共参数
		c.Set("example", "12345")
		//等其他中间件先执行
		c.Next()
		//获取耗时
		latency := time.Since(t)
		log.Print("接口耗时时间：", latency)
	}
}

/**
 * 渲染html页面
 */
func html(c *gin.Context) {
	username := c.Query("username")
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是gin", "name": username})
}

/**
 *  返回json对象
 */
func json(c *gin.Context) {
	var user = new(models.UserModel)
	user.Name = "Lena"
	user.Message = "hey"
	user.Number = 123
	c.JSON(200, user)
}
