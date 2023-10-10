package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var MyMiddleWare = func(c *gin.Context) {
	//请求唯一id，用于追踪链路上的异常模块，一般体现在日志中
	uuidValue := uuid.New()
	c.Set("uuid", uuidValue)
}
