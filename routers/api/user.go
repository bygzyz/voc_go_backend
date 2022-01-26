package api

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/request"
	"github.com/EDDYCJY/go-gin-example/pkg/response"
	"github.com/EDDYCJY/go-gin-example/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUser 获取用户
func GetUser(c *gin.Context) {

	appG := app.Gin{C: c}
	// 绑定参数
	var req request.UserRequestStruct
	err := c.ShouldBind(&req)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	users, err := models.GetUser(&req)

	// 转为ResponseStruct, 隐藏部分字段
	//var respStruct []response.UserListResponseStruct
	var respStruct []response.UserListResponseStruct
	utils.Struct2StructByJson(users, &respStruct)
	// 返回分页数据
	var resp response.PageData
	// 设置分页参数
	//resp.PageInfo = req.PageInfo
	// 设置数据列表
	resp.List = respStruct
	response.SuccessWithData(resp)

	//appG.Response(http.StatusOK, e.SUCCESS, resp)
	return
	//if err != nil {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
	//	return
	//} else {
	//	c.JSON(http.StatusBadRequest, gin.H{"userList": users})
	//	return
	//}
}

// CreateUser 创建用户
func CreateUser(c *gin.Context) {

	appG := app.Gin{C: c}
	var createUser request.CreateUserRequestStruct
	if err := c.ShouldBindJSON(&createUser); err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	username := createUser.Username
	password := createUser.Password

	userId, err := models.CreateUser(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	} else {
		appG.Response(http.StatusCreated, e.SUCCESS, map[string]uint{"userId": userId})
		return
	}
}

// UpdateUser 修改用户
func UpdateUser(c *gin.Context) {
	appG := app.Gin{C: c}
	var updateUser request.UpdateUserRequestStruct
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	userId := updateUser.UserId
	username := updateUser.Username

	_, err := models.UpdateUser(userId, username)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	} else {
		appG.Response(http.StatusAccepted, e.SUCCESS, nil)
		return
	}
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	// TODO 根据用户id删除指定用户
	appG := app.Gin{C: c}
	var deleteUser request.DeleteUserRequestStruct
	if err := c.ShouldBindJSON(&deleteUser); err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	userId := deleteUser.UserId

	_, err := models.DeleteUser(userId)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	} else {
		appG.Response(http.StatusOK, e.SUCCESS, nil)
		return
	}
}
