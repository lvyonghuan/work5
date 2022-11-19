package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var username string
var password string

func getcookie(c *gin.Context) {
	_, err := c.Request.Cookie("login_cookie")
	if err != nil {
		c.Abort()
	}
	fmt.Println("get successful")
}

func login(c *gin.Context) {
	username = c.PostForm("username")
	password = c.PostForm("password")
	fmt.Println("reveive:", username, password)
	login_cookie := &http.Cookie{
		Name:     "login_cookie",
		Value:    username,
		Path:     "/",
		HttpOnly: false,
		MaxAge:   1919,
	}
	c.Writer.Header().Set("Set-Cookie", login_cookie.String())
}

func main() {
	r := gin.Default()
	usr := r.Group("/user")
	{
		usr.POST("/login", login, getcookie, func(c *gin.Context) {
			c.JSON(200, username)
			c.JSON(200, password)
		})
	}
	r.Run()
}

//这下真整不明白了（指怎么检测登陆状态保持
