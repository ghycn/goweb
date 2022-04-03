package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS int = 1000
	FAIL    int = 2000

	SucceedMsg string = "操作成功！"
	FailedMsg  string = "操作失败！"
)

// 消息实体
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

// 成功
func Success(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, SucceedMsg, c)
}

// 成功，自定义提示信息
func SuccessMsg(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

// 失败
func Fail(c *gin.Context) {
	Result(FAIL, map[string]interface{}{}, FailedMsg, c)
}

// 失败，自定义提示信息
func FailMsg(message string, c *gin.Context) {
	Result(FAIL, map[string]interface{}{}, message, c)
}
