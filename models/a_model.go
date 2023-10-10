package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type AModel struct {
	ID  string `json:"id"`
	MSG string `json:"msg"`
}

// 结构体标准输出函数，一般不需要自定义
func (this *AModel) String() string {
	return fmt.Sprintf("id: %s, msg: %s", this.ID, this.MSG)
}

func InitA(c *gin.Context) *AModel {
	uuid, _ := c.Get("uuid")
	return &AModel{ID: fmt.Sprintf("%s", uuid), MSG: "get a"}
}
