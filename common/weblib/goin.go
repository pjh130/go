package weblib

//Gin is a web framework written in Golang
import "github.com/gin-gonic/gin"

func ExampleGoin() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
