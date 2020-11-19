package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("------------")

	router := gin.Default()
	router.Static("/static", "./static")

	router.GET("/index", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "ok",
		})
	})

	router.Run(":7070")

}
