package response

import (
	"github.com/gin-gonic/gin"
	"github.com/liuxianbin/pkg/check"
	"github.com/liuxianbin/pkg/errcode"
	"log"
)

func ReturnJson(c *gin.Context, httpCode int, dataCode int, msg string) {
	c.JSON(httpCode, gin.H{
		"code":    dataCode,
		"message": msg,
	})
}
func Success(c *gin.Context) {
	ReturnJson(c, errcode.OK.HttpCode(), errcode.OK.Code(), errcode.OK.Message())
}

func Fail(c *gin.Context, err *errcode.Error) {
	ReturnJson(c, err.HttpCode(), err.Code(), err.Message())
}

func Failx(c *gin.Context, err *errcode.Error, appendMsg string) {
	ReturnJson(c, err.HttpCode(), err.Code(), err.Message()+": "+appendMsg)
}

// Auto 增删改响应结果
func Auto(c *gin.Context, err error) {
	if err != nil {
		Fail(c, errcode.InternalServerError)
	} else {
		Success(c)
	}
}

// BindRequest 绑定校验请求参数
func BindRequest(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBind(obj); err != nil {
		log.Println(err)
		Failx(c, errcode.BadRequest, "json数据格式错误")
		return false
	}
	return true
}

// Check 校验必填项
func Check(c *gin.Context, obj interface{}) bool {
	if err := check.Do(obj); err != nil {
		Failx(c, errcode.BadRequest, err.Error())
		return false
	}
	return true
}

// BindRequestAndCheck 绑定校验请求参数并校验
func BindRequestAndCheck(c *gin.Context, obj interface{}) bool {
	if !BindRequest(c, obj) {
		return false
	}
	return Check(c, obj)
}
