package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
	h "github.com/lecec/lecec/handler"
)

// Default 功能实现
func Default() (r *gin.Engine) {
	r = gin.Default()
	health := &h.Health{}
	r.GET("/health", func(c *gin.Context) {
		health.Health(c)
	})
	r.POST("/rpc/:service/:method", func(c *gin.Context) {
		fmt.Println(c, c.PostForm, c.Param("service"), c.Param("method"), c.Params)
		var json map[string]interface{}
		c.ShouldBindJSON(&json)
		fmt.Println(json)
		c.JSON(200, gin.H{
			"Valid":   true,
			"service": c.Param("service"),
			"method":  c.Param("method"),
		})
	})
	return r
}
