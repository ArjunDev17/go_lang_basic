package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("API")
	r := gin.Default()
	r.GET("/get", getValues)
	r.POST("/post", postValues)
	r.PUT("/put", putValues)
	r.DELETE("/delete", deleteValues)
	r.Run()
}
