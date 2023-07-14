package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/liuxianbin/pkg/check"
	"github.com/liuxianbin/pkg/errorx"
)

func ReturnJson(c *gin.Context, httpCode int, dataCode int, msg string) {
	c.JSON(httpCode, gin.H{
		"code":    dataCode,
		"message": msg,
	})
}
func Success(c *gin.Context) {
	ReturnJson(c, errorx.OK.HttpCode(), errorx.OK.Code(), errorx.OK.Message())
}

func Fail(c *gin.Context, err *errorx.Error) {
	ReturnJson(c, err.HttpCode(), err.Code(), err.Message())
}

func Failx(c *gin.Context, err *errorx.Error, appendMsg string) {
	ReturnJson(c, err.HttpCode(), err.Code(), err.Message()+": "+appendMsg)
}

// Auto 增删改响应结果
func Auto(c *gin.Context, err error) {
	if err != nil {
		Fail(c, errorx.InternalServerError)
	} else {
		Success(c)
	}
}

// BindRequest 绑定校验请求参数
func BindRequest(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBind(obj); err != nil {
		Failx(c, errorx.BadRequest, "json数据格式错误")
		return false
	}
	return true
}

// Check 校验必填项
func Check(c *gin.Context, obj any) bool {
	if err := check.Do(obj); err != nil {
		Failx(c, errorx.BadRequest, err.Error())
		return false
	}
	return true
}

// BindRequestAndCheck 绑定校验请求参数并校验
func BindRequestAndCheck(c *gin.Context, obj any) bool {
	if !BindRequest(c, obj) {
		return false
	}
	return Check(c, obj)
}

func BindRequestWithTranslates(c *gin.Context, obj any) bool {
	if err := c.ShouldBind(obj); err != nil {
		if errs, ok := TranslateErrorWithContext(c, err); ok {
			for _, v := range errs {
				Failx(c, errorx.BadRequest, v)
				break
			}
			return false
		}
		Failx(c, errorx.BadRequest, "json数据格式错误")
		return false
	}
	return true
}
