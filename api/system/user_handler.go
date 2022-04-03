package system

import (
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

// 删除用户
func Delete(c *gin.Context) {
	user := models.User{}
	id := c.Param("id")
	// 获取全部记录
	affected := sqlgorm.JdbcTemplate.Model(user).Where("id = ?", id).Delete(&user).RowsAffected
	if affected == 1 {
		result.Success(affected, c)
	} else {
		result.Fail(c)
	}
}

// 保存或者修改用户
func SaveOrUpdate(c *gin.Context) {
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		result.FailMsg("user对象json解析异常:"+err.Error(), c)
		global.GVA_LOG.Error("user对象json解析异常", zap.Error(err))
		return
	}
	//var affected int64
	if &user.ID == nil || user.ID == 0 {
		// 保存用户
		sqlgorm.JdbcTemplate.Model(user).Save(&user)
	} else {
		// 修改用户
		sqlgorm.JdbcTemplate.Model(user).Update(&user)
	}
	result.Success(&user.ID, c)
}
