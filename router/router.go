package router

import (
	userApi "gin-web/api/system"
	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter(model string, path string) *gin.Engine {
	gin.SetMode(model)
	r := gin.Default()
	group := r.Group(path)
	// 文章相关的接口组
	router(group)
	return r
}

func router(group *gin.RouterGroup) {
	// 用户相关的接口
	userRouter := group.Group("/user")
	{
		// 获取用户列表
		userRouter.GET("/list", userApi.List)
		// 删除用户
		userRouter.DELETE("/:id", userApi.Delete)
		// 添加用户
		userRouter.POST("/", userApi.Add)
	}
}
