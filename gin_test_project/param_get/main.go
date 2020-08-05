package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		FirstName := c.Query("first_name")
		LastName := c.DefaultQuery("last_name", "last_name_value")
		c.String(http.StatusOK, "%s,%s", FirstName, LastName)
	})

	r.Run() // 8080
}
