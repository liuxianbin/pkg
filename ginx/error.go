package ginx

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// TranslateError 把error翻译成中文
func TranslateError(trans ut.Translator, err error) (validator.ValidationErrorsTranslations, bool) {
	vErrors, ok := err.(validator.ValidationErrors)
	if ok {
		errs := vErrors.Translate(trans)
		return errs, true
	}
	return nil, false
}

// TranslateErrorWithContext 从Gin上下文中获取Translator翻译error成中文
func TranslateErrorWithContext(c *gin.Context, err error) (validator.ValidationErrorsTranslations, bool) {
	trans, ok := c.Get("trans")
	if ok {
		trans := trans.(ut.Translator)
		return TranslateError(trans, err)
	}
	return nil, false
}
