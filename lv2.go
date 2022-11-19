package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var use_pass = make(map[string]string) //储存用户信息

func register(c *gin.Context) { //注册段
	var username string
	var password string
	username = c.PostForm("username")
	password = c.PostForm("password")
	use_pass[username] = password
	fmt.Println("注册成功。用户名:" + username + ",密码：" + password)
	c.JSON(200, "注册成功！")
}

func login_1(c *gin.Context) { //登录段
	var username string
	var password string
	username = c.PostForm("username")
	password = c.PostForm("password")
	value, ok := use_pass[username] //登录检查
	if ok == false {
		c.JSON(200, "用户未注册")
		return
	}
	if value != password {
		c.JSON(200, "密码错误！")
		return
	}
	fmt.Println("登陆成功。用户名:" + username + ",密码：" + password)
	c.JSON(200, "登陆成功！")
}

func main() {
	r := gin.Default()
	r.GET("/register", register)
	r.GET("/login", login_1)
	r.Run()
}
