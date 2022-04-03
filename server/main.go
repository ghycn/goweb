package main

import (
	"fmt"
	"gin-web/common/global"
	"gin-web/config"
	"gin-web/database/sqlgorm"
	"gin-web/router"
	"gin-web/utils"
)

func main() {
	fmt.Println("Gin-Web，启动")
	// 读取系统配置
	sysConfig := config.NewConfig().ReadConfig()
	// 初始化zap日志库
	global.GVA_LOG = utils.Zap()
	// 加载Banner
	config.ReadBanner(sysConfig.Server.Banner.Name)
	// 加载数据库连接
	sqlgorm.InitDb(sysConfig)
	router.InitRouter(sysConfig.Server.Model, sysConfig.Server.Path).Run(sysConfig.Server.Part)

	//// 常规的初始化路由
	//router := gin.Default()
	//router.GET("/", func(c *gin.Context) {
	//time.Sleep(5 * time.Second)
	//	c.String(http.StatusOK, "Welcome Gin Server")
	//})
	//// 定义服务器
	//srv := &http.Server{
	//	Addr:    ":8080",
	//	Handler: router,
	//	}
	//	// 利用 goroutine 启动监听
	//	go func() {
	//		// srv.ListenAndServe() 监听
	//		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//			log.Fatalf("listen: %s\n", err)
	//		}
	//	}()
	//	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	//	quit := make(chan os.Signal)
	//	signal.Notify(quit, os.Interrupt)
	//	// quit 信道是同步信道，若没有信号进来，处于阻塞状态
	//	// 反之，则执行后续代码
	//	<-quit
	//	log.Println("Shutdown Server ...")
	//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//	defer cancel()
	//	// 调用 srv.Shutdown() 完成优雅停止
	//	// 调用时传递了一个上下文对象，对象中定义了超时时间
	//	if err := srv.Shutdown(ctx); err != nil {
	//		log.Fatal("Server Shutdown:", err)
	//	}
	//	log.Println("Server exiting")

}
