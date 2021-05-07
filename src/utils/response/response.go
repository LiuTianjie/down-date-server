package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Reason string      `json:"reason"`
}

func Result(code int, data interface{}, msg string, reason string, c *gin.Context) {
	// 开始时间
	c.JSON(code, Response{
		code,
		data,
		msg,
		reason,
	})
}

func Ok(code int, c *gin.Context) {
	Result(code, map[string]interface{}{}, "SUCCESS", "操作成功", c)
}

func OkWithMessage(code int, reason string, c *gin.Context) {
	Result(code, map[string]interface{}{}, "SUCCESS", reason, c)
}

func OkWithDetailed(code int, data interface{}, reason string, c *gin.Context) {
	Result(code, data, "SUCCESS", reason, c)
}

func Fail(code int, c *gin.Context) {
	Result(code, map[string]interface{}{}, "FAILED", "操作失败", c)
}

func FailWithMessage(code int, reason string, c *gin.Context) {
	Result(code, map[string]interface{}{}, "FAILED", reason, c)
}

func FailWithDetailed(code int, data interface{}, reason string, c *gin.Context) {
	Result(code, data, "FAILED", reason, c)
}
