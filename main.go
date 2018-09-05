package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/demos/:id", GetDemoHandler)
	r.POST("/demos", PostDemoHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}
