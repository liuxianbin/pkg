package errorx

import (
	"fmt"
	"net/http"
)

func ExampleNewError() {
	e := NewError(http.StatusBadRequest, 100400, "参数错误")
	fmt.Println(e.Message(), e.Code())
	fmt.Println(OK.Code(), NotFound.Code(), NotFound.HttpCode())
	// output:
	// 参数错误 100400
	// 0 100002 404
}
