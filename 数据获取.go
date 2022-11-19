package main

//	func main() {
//		r := gin.Default()
//		r.GET("/ping", func(c *gin.Context) {
//			value := c.DefaultQuery("key", "null")
//			val = c.Query("key")
//		})
//		r.Run()
//	}
//
//	func main() {
//		r := gin.Default()
//		r.GET("/ee", func(c *gin.Context) {
//			c.JSON(200, "这太谔谔了")
//		})
//		r.POST("/ee", func(c *gin.Context) {
//			c.JSON(200, "你说得对，但是这很谔谔")
//		})
//		r.Run()
//	}
//
//	func main() {
//		r := gin.Default()
//		r.GET("/hao", func(c *gin.Context) {
//			c.JSON(200, c.Query("hao"))
//		})
//		r.Run()
//	}

//type Login struct {
//	Username string `json:"username,omitempty" form:"username"`
//	Password string `json:"password,omitempty" form:"password"`
//}
//
//func main() {
//	r := gin.Default()
//	r.POST("/ping", func(c *gin.Context) {
//		var login Login
//		if err := c.ShouldBind(&login); err != nil {
//			log.Println(err)
//			c.JSON(400, "err")
//			return
//		}
//		c.JSON(200, login)
//		return
//	})
//	r.Run()
//}

//func main() {
//	r := gin.Default()
//	r.POST("/ping", func(c *gin.Context) {
//		file, _ := c.FormFile("file")
//		log.Println(file.Filename)
//		dst := "./" + file.Filename
//		c.SaveUploadedFile(file, dst)
//		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
//	})
//	r.Run()
//}
