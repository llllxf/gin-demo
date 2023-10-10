package routers

import (
	"gin-demo/middleware"
	"gin-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {

}

type AHandler struct{}

func ARouterGroup(router *gin.Engine) {
	aHandler := &AHandler{}

	group := router.Group("/a")
	group.Use(middleware.MyMiddleWare) //给 a 路由组所有路由配置中间件
	{
		group.GET("/get", aHandler.Get)
	}
}

func (*AHandler) Get(c *gin.Context) {

	a := service.GetA(c)
	c.JSON(http.StatusOK, a)

}
