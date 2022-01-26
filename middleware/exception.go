package middleware

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 全局异常处理中间件
func Exception(c *gin.Context) {
	fmt.Println("进入了全局异常处理函数内")
	defer func() {
		if err := recover(); err != nil {
			if resp, ok := err.(response.Resp); ok {
				fmt.Println(err)
				resp.RequestId = "1"
				// 以json方式写入响应
				response.JSON(c, http.StatusBadRequest, resp)
			} else {
				resp := response.Resp{
					Code:      response.InternalServerError,
					Data:      map[string]interface{}{},
					Msg:       response.CustomError[response.InternalServerError],
					RequestId: resp.RequestId,
				}
				// 以json方式写入响应
				response.JSON(c, http.StatusOK, resp)
				c.Abort()
				return
			}
		}
	}()
	c.Next()
}
