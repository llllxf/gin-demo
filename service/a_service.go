package service

import (
	"gin-demo/funcs"
	"gin-demo/models"
	"github.com/gin-gonic/gin"
)

func GetA(c *gin.Context) *models.AModel {

	return funcs.GetA(c)

}
