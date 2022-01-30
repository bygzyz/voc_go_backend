package api

import (
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/request"
	"github.com/EDDYCJY/go-gin-example/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DataExploreFilterMessage 数据探索-信息列表页
func DataExploreFilterMessage(c *gin.Context) {

	appG := app.Gin{C: c}
	// 绑定参数
	var req request.DataExploreRequestStruct
	err := c.ShouldBind(&req)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	//es, _ := elasticsearch.NewDefaultClient()
	//log.Println(elasticsearch.Version)
	//log.Println(es.Info())

	var respStruct []response.UserListResponseStruct
	//utils.Struct2StructByJson(users, &respStruct)
	// 返回分页数据
	var resp response.PageData
	// 设置分页参数
	// 设置数据列表
	resp.List = respStruct
	response.SuccessWithData(resp)

	//appG.Response(http.StatusOK, e.SUCCESS, resp)
	return

}
