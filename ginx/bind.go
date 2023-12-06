package ginx

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func MustInt(s string) int {
	id, _ := strconv.Atoi(s)
	return id
}

func MustQueryInt(c *gin.Context, key string) int {
	return MustInt(c.Query(key))
}

func MustParamInt(c *gin.Context, key string) int {
	return MustInt(c.Param(key))
}

func MustParamInts(c *gin.Context, key string) []int {
	var arr []int
	vs := strings.Split(c.Param(key), ",")
	for _, v := range vs {
		if vv, err := strconv.Atoi(v); err == nil {
			arr = append(arr, vv)
		}
	}
	return arr
}
