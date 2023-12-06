package ginx

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
	"github.com/juju/ratelimit"
	"github.com/liuxianbin/pkg/errorx"
	"github.com/liuxianbin/pkg/logx"
	"log"
	"os"
	"time"
)
import "github.com/gin-gonic/gin"

var Logger = logx.New(os.Stdout, "gin_log: ", log.LstdFlags)

// ContextTimeout Context超时控制
func ContextTimeout(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// Recovery 异常恢复
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				Fail(c, errorx.InternalServerError)
				Logger.WithCallers().Json("recover")
				c.Abort()
			}
		}()
		c.Next()
	}
}

type multiWrite struct {
	gin.ResponseWriter
	buf bytes.Buffer
}

func (w *multiWrite) write(p []byte) (int, error) {
	if n, err := w.buf.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

// AccessLog 访问日志
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := &multiWrite{
			ResponseWriter: c.Writer,
			buf:            bytes.Buffer{},
		}
		c.Writer = w
		c.Next()
		fields := logx.Fields{
			"url":    c.Request.URL.Path,
			"method": c.Request.Method,
			"status": w.Status(),
			"param":  c.Request.PostForm.Encode(),
			"body":   w.buf.String(),
		}
		Logger.WithFields(fields).Json("access")
	}
}

// TranslateZH Gin校验信息翻译成中文
func TranslateZH() gin.HandlerFunc {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		zh2.RegisterDefaultTranslations(v, trans)
	}
	return func(c *gin.Context) {
		c.Set("trans", trans)
		c.Next()
	}
}

// RateLimiter 限流
func RateLimiter(fillInterval time.Duration, capacity, quantum int64) gin.HandlerFunc {
	b := ratelimit.NewBucketWithQuantum(fillInterval, capacity, quantum)
	return func(c *gin.Context) {
		if b.TakeAvailable(1) == 0 {
			Fail(c, errorx.TooManyRequests)
			c.Abort()
			return
		}
		c.Next()
	}
}
