package gin

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	h "github.com/lecec/lecec/handler"
	"google.golang.org/grpc"

	pb "github.com/lecec/userApi/proto/health"
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

	r.GET("/userapi", func(c *gin.Context) {
		fmt.Println(c)
		// Set up a connection to the server.
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		client := pb.NewHealthClient(conn)
		req := &pb.Request{}
		res, err := client.Health(c, req)
		log.Println(res, err)
	})
	return r
}
