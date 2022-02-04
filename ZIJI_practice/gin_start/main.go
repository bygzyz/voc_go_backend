package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由和路由分组的基本使用
// gin router

func main() {
	r := gin.Default()
	// gin default adds default middleware just like
	//Logger which same as stdout
	//and
	//Recovery which return 500 when program panic
	//r.GET("/goodsList",getGoods)
	newGroup := r.Group("/goods")
	{
		newGroup.GET("/list", getGoods)
		newGroup.POST("/add", addGoods)
	}

	r.Run(":8081")

}

func addGoods(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "add success",
	})

}

func getGoods(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"message": "gin practise",
	})

}
