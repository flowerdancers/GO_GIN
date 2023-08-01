package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

//create a new random string
func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterslength := len(letters)
	result := ""
	for i := 0; i < n; i++ {
		r := letters[rand.Intn(letterslength)]
    result += string(r)
	}
	return result
}
func main() {
	r := gin.Default()
	r.GET("/api/auth/register", func(c *gin.Context) {
		//获取参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")

		//数据验证
		if len(telephone) != 11 {
			c.JSON(400, gin.H{"code": 400, "msg": "The mobile phone number must contain 11 characters "})
      return
		}
		if len(password) < 6 {
      c.JSON(400, gin.H{"code": 400, "msg": "The password must contain at least 6 characters "})
      return
    }
		//if no name ,give a random string of 11 digits
		if len(name) == 0 {
      name = RandomString(10)
      return
    }
    //判断手机号是否存在
    if isExist(telephone) {
      c.JSON(400, gin.H{"code": 400, "msg": "The mobile phone number has been registered "})
      return
    }
    //保存用户信息
    saveUser(name, telephone, password)
    c.JSON(200, gin.H{"code": 200, "msg": "Register success"})

	})

	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}
