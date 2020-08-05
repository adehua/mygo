package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		bodyByts, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()

		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByts))
		FirstName := c.PostForm("first_name")
		LastName := c.DefaultPostForm("last_name", "last_name_value")
		c.String(http.StatusOK, "%s,%s,%s", FirstName, LastName, string(bodyByts))
	})

	r.Run() // 8080
}
