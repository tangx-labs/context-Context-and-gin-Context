package main

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_GinContext(t *testing.T) {

	r := gin.Default()

	r.GET("/", handler)
	_ = r.Run()

}

func handler(c *gin.Context) {
	// c.Value("key")
	// c.Query("eky")
	c.Set("key", 123)
	ret := c.Query("key")
	fmt.Println("ret= ", ret)
}
