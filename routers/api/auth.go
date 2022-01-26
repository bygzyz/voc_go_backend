package api

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `valid:"Required; MaxSize(50)" json:"username"`
	Password string `valid:"Required; MaxSize(50)" json:"password"`
}

// @Summary Get Auth
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func Login(c *gin.Context) {

	appG := app.Gin{C: c}
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	username := user.Username
	password := user.Password

	isExist, err := models.CheckAuth(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_USER_PASSWORD_ERROR, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_USER_PASSWORD_ERROR, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

// Logout 退出登录
func Logout(c *gin.Context) {

}
