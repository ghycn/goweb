package system

import (
	"fmt"
	"gin-web/common/global"
	"gin-web/common/result"
	"gin-web/database/sqlgorm"
	"gin-web/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 查询用户列表
func List(c *gin.Context) {

	var users []models.User
	// 获取全部记录
	sqlgorm.JdbcTemplate.Find(&users)
	result.Success(&users, c)
}

// 查询用户列表
func Delete(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	// 获取全部记录
	affected := sqlgorm.JdbcTemplate.Where("id = ?", id).Delete(&models.User{}).RowsAffected
	result.Success(affected, c)
}

// 查询用户列表
func Add(c *gin.Context) {
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		result.FailMsg("user对象json解析异常:"+err.Error(), c)
		global.GVA_LOG.Error("user对象json解析异常", zap.Error(err))
		return
	}
	// 保存用户
	affected := sqlgorm.JdbcTemplate.Save(&user).RowsAffected
	result.Success(affected, c)
}
