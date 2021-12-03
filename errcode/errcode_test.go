package errcode

import (
	"net/http"
)

func ExampleNewError() {
	NewError(http.StatusBadRequest, 100400, "参数错误")
	//output:
	//
}
