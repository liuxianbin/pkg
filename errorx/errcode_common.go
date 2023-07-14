package errorx

import "net/http"

var (
	OK                  = NewError(http.StatusOK, 0, "成功")
	BadRequest          = NewError(http.StatusBadRequest, 100001, "入参错误")
	NotFound            = NewError(http.StatusNotFound, 100002, "资源不存在")
	InternalServerError = NewError(http.StatusInternalServerError, 100003, "系统异常")
	Unauthorized        = NewError(http.StatusUnauthorized, 100004, "鉴权失败")
)
