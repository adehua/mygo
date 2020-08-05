package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"github.com/go-playground/validator/v10"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_in" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", func(c *gin.Context) {
		var b Booking
		if err := c.ShouldBind(b); err != nil {
			c.String(500, "%v", err.Error())
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "ok！", "booking": b})

	})

	r.Run() // 8080
}
