package main

import (
	"github.com/EDDYCJY/go-gin-example/ZIJI_practice/gin_path_params/proto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin  路径参数获取和基本注意事项
// 使用路径参数也可以通过结构体严格校验路径参数类型
// query 和 form表单 数据解析

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	r := gin.Default()
	// 路由分组，组内路径会自动包含前缀
	group := r.Group("goods")
	{
		//group.GET("",goodsList)
		group.GET("/:id/:name", goodsDetail)
		group.GET("/user", userList)
		group.GET("/someProtoBuf", someProtoBuf)
		group.POST("/user", addUser)
	}

	r.Run(":8082")

}

func someProtoBuf(c *gin.Context) {
	course := []string{"go", "python", "微服务"}
	userInfo := &proto.Teacher{
		Name:   "bobby",
		Course: course,
	}
	c.ProtoBuf(http.StatusOK, userInfo)

}

func addUser(context *gin.Context) {
	userId := context.PostForm("id")
	userName := context.DefaultPostForm("username", "bobby")
	context.JSON(http.StatusOK, gin.H{
		"userid":   userId,
		"username": userName,
	})

}

func userList(c *gin.Context) {
	userid := c.Query("id")
	userName := c.DefaultQuery("name", "bobby")

	var person struct {
		Name string `json:"name"`
		Age  string `json:"age"`
	}
	person.Name = userName
	person.Age = userid

	c.JSON(http.StatusOK, person)

}

func goodsDetail(c *gin.Context) {
	var person Person
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   person.ID,
		"name": person.Name,
	})

}

func goodsList(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
