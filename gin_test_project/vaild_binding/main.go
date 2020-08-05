package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Age     int `form:"age" binding:"required,lt=10"`
	Name    string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/testing",func (c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, "%v", err)
			c.Abort()
			return
		}
		c.String(200, "%v", person)

	})


	r.Run() // 8080
}

