package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/login", login)
		v1.GET("submit", submit)
	}

	v2 := r.Group("/v2")
	{
		v2.POST("login", login)
		v2.POST("submit", submit)
	}

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "hehe")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})

	r.Run(":8000")
}

func login(c *gin.Context)  {
	name := c.DefaultQuery("name", "jack")
	//c.String(http.StatusOK, fmt.Sprintf("hello %s\n", name))
	c.JSON(http.StatusOK, gin.H{"name":name})
}

func submit(c *gin.Context)  {
	name := c.DefaultQuery("name", "lily")
	c.String(http.StatusOK, fmt.Sprintf("hello %s\n", name))
}
