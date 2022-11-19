package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/regiester", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "你说的都对"
			c.SetCookie("gin_cokie", "henhenhenaaaaaa", 1919, "/", "localhost", false, true)
		}
		fmt.Println("cokkie value:%s\n", cookie)
	})
	r.Run()
}

//func setCookieHandler(c *gin.Context) {
//	test_cookie := &http.Cookie{
//		Name:     "test_cookie",
//		Value:    "username",
//		Path:     "/",
//		HttpOnly: false,
//		MaxAge:   3600,
//	}
//	c.Writer.Header().Set("Set-Cookie", test_cookie.String())
//	//http.SetCookie(c.Writer,test_cookie)
//}
