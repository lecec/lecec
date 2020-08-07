package handler

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Health 用户结构
type Health struct {
}

// Health 用户是否存在
func (srv *Health) Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"Valid": true,
		"Time":  time.Now().Format("2006-01-02 15:04:05"),
	})
}
