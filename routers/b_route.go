package routers

import (
	"gin-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {

}

type BHandler struct{}

func BRouterGroup(router *gin.Engine) {
	bHandler := &BHandler{}

	group := router.Group("/b")
	{
		group.POST("/get", bHandler.Get)
	}
}

func (*BHandler) Get(c *gin.Context) {

	param := c.PostForm("param")

	if param == "" {
		c.JSON(http.StatusOK, "lack param")
		return
	}
	b := service.GetB(c, param)

	c.JSON(http.StatusOK, b)

}
