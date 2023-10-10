package service

import (
	"gin-demo/funcs"
	"gin-demo/models"
	"github.com/gin-gonic/gin"
)

func GetB(c *gin.Context, param string) *models.BModel {

	return funcs.GetB(c, param)

}
