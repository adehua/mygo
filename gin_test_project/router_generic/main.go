package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/user/*action", func(c *gin.Context) {
		c.String(200, "action")
	})

	r.Run() // 8080
}
