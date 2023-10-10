package funcs

import (
	"gin-demo/models"
	"github.com/gin-gonic/gin"
)

func GetA(c *gin.Context) *models.AModel {

	a := models.InitA(c)
	return a

}
