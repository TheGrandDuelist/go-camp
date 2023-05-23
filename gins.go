package main

import (
	"fmt"
	gin "github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type LoginModel struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func entry_9() {
	r := gin.Default()
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		//截取/
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})
	r.GET("/user", func(context *gin.Context) {
		name := context.DefaultQuery("name", "test1")
		context.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	v1 := r.Group("/v1")
	{
		v1.POST("/login", login)
		v1.GET("/submit", submit)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/post", post)
		v2.GET("/del", del)
	}
	//默认为监听8080端口
	r.Run(":8000")
}

func del(context *gin.Context) {
	context.String(http.StatusOK, fmt.Sprintf("hello %s", "del"))
}

func post(context *gin.Context) {
	context.String(http.StatusOK, fmt.Sprintf("hello %s", "post"))
}

func submit(context *gin.Context) {
	name := context.DefaultQuery("name", "jack")
	context.String(http.StatusOK, fmt.Sprintf("submit %s", name))
}

func login(c *gin.Context) {
	var loginModel LoginModel
	if err := c.ShouldBindJSON(&loginModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 200, "data": loginModel})
}
