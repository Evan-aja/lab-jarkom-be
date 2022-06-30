package main

import (
	database "lab-jarkom-be/Database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//Database
	db := database.Open()
	if db != nil {
		println("nice")
	}
	r := gin.Default()

	//CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Controll-Allow-Methods", "POST,OPTIONS,GET,PUT,DELETE,PATCH")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	//Route taruh disini
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Lab-JarKom",
			"success": true,
		})
	})

	if err := r.Run(":5000"); err != nil {
		println("Error")
		println(err.Error())
		return
	}
	println("Server jalan tuh")
}
