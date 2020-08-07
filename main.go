package main

import (
	// 公共引入
	"github.com/lecec/lecec/gin"
)

func main() {
	r := gin.Default() // 路由实现
	r.Run(":8080")     // listen and serve on 0.0.0.0:8080
}
