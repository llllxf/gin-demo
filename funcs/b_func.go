package funcs

import (
	"gin-demo/models"
	"github.com/gin-gonic/gin"
)

func GetB(c *gin.Context, param string) *models.BModel {

	b := models.InitB(param)
	return b

}
