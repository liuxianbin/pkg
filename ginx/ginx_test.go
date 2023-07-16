package ginx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

type Info struct {
	Content string `json:"content" binding:"required"`
	Num     int    `json:"num" binding:"required"`
}

func TestGinx(t *testing.T) {
	e := gin.New()
	e.Use(AccessLog())
	e.Use(Recovery())
	e.Use(ContextTimeout(time.Second))
	e.Use(TranslateZH())

	e.GET("/demo", func(c *gin.Context) {
		c.String(http.StatusOK, "demo...")
		log.Println("demo...")
		time.Sleep(time.Second * 3)
		panic("exception...")
	})

	e.POST("/demo2", func(c *gin.Context) {
		log.Println("demo2...")
		var info Info
		if !BindRequestWithTranslates(c, &info) {
			return
		}
		c.String(http.StatusOK, "demo2...")
		fmt.Println("ok")
	})

	param, _ := json.Marshal(map[string]any{
		"content": "default",
		//"num":     20,
	})
	params := bytes.NewReader(param)
	req := httptest.NewRequest("POST", "/demo2", params)
	req2 := httptest.NewRequest("GET", "/demo", nil)
	req.Header.Set("content-type", "application/json")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	result := w.Body.String()
	fmt.Println(result)
	e.ServeHTTP(w, req2)
}

func TestRateLimiter(t *testing.T) {
	e := gin.Default()
	e.GET("/demo3", RateLimiter(time.Second, 10, 10), func(c *gin.Context) {
		log.Println("demo3...")
		c.String(http.StatusOK, "demo3...")
	})
	req := httptest.NewRequest("GET", "/demo", nil)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		w := httptest.NewRecorder()
		wg.Add(1)
		go func() {
			defer wg.Done()
			e.ServeHTTP(w, req)
			result := w.Body.String()
			fmt.Println(result)
		}()
	}
	wg.Wait()
	fmt.Println("over")
}
