package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//自定义中间件

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("key", "init")
		timeUsed := time.Since(t)
		fmt.Printf("使用的时间是 %s", timeUsed)

		status := c.Writer.Status()
		fmt.Println("当前的状态", status)
		c.Next()

	}
}

func main() {

	router := gin.New()
	//router.Use(gin.Logger(),gin.Recovery())
	router.Use(MyLogger())

	router.GET("/goods", goodsList)

	// 优雅退出
	go func() {
		router.Run(":8082")
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGALRM)
	<-quit

	//处理后续逻辑
	fmt.Println("关闭server中")

}

func goodsList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
	return
}
