package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.StaticFS("/stati", http.Dir("static"))
	r.StaticFile("./favicon.ico", "./favicon.ico")

	r.Run() // 8080
}
