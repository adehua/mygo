package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Person struct {
	Name     string    `from:"name"`
	Address  string    `from:"address"`
	Birthday time.Time `from:"birthday"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", testing)
	r.POST("/testing", testing)

	r.Run() // 8080
}

func testing(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err == nil {
		c.String(200, "%v", person)
	}

	//c.String(200, "person bind error", err)

}
