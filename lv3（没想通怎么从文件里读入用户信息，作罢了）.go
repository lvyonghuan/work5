//package main
//
//import (
//	"github.com/gin-gonic/gin"
//	"os"
//)
//
//var userinfo *os.File
//
//func newregister(c *gin.Context) {
//	var (
//		username string
//		password string
//	)
//	username = c.PostForm("username")
//	password = c.PostForm("password")
//
//}
//
//func main() {
//	_, err := os.Stat("./userinfo.txt")
//	if err != nil {
//		userinfo, _ = os.Create("./userinfo.txt")
//	}
//	r := gin.Default()
//	r.GET("/register", newregister)
//}
