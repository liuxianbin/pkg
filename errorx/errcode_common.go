package errorx

import "net/http"

var (
	OK                  = NewError(http.StatusOK, 0, "成功")
	BadRequest          = NewError(http.StatusBadRequest, 100001, "入参错误")
	NotFound            = NewError(http.StatusNotFound, 100002, "资源不存在")
	InternalServerError = NewError(http.StatusInternalServerError, 100003, "系统异常")
	Unauthorized        = NewError(http.StatusUnauthorized, 100004, "认证失败")
	Forbidden           = NewError(http.StatusForbidden, 100005, "鉴权失败")
	TooManyRequests     = NewError(http.StatusTooManyRequests, 100006, "Too Many Requests")
)
