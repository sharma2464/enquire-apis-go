package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/services"
	"main/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/replit/database-go"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		// c.Header("Access-Control-Allow-Credentials", "true")
		// c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "*")
		// c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		// c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
		c.Header("Access-Control-Allow-Methods", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	fmt.Println("Hello, World!")
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Health Check Done",
		})
	})

	r.POST("/enquire", func(c *gin.Context) {
		var form types.EnquireForm
		if c.ShouldBind(&form) == nil {
			if form.FirstName != "" && form.LastName != "" && form.PhoneNumber != "" && form.Email != "" && form.NumberOfPets != "" && form.AnimalType != "" && form.Pincode != "" {
				b, err := json.Marshal(form)
				if err != nil {
					resErr, _ := fmt.Printf("Error: %s", err)
					c.JSON(500, gin.H{"status": resErr})
					return
				}

				dt := database.Set(string(form.Email), string(b))
				if dt != nil {
					c.JSON(http.StatusOK, gin.H{"status": "Enquiry Created, Someone will get in touch with you soon.", "data": dt})
				} else {
					c.JSON(500, gin.H{"status": "Unable to create Enquiry after payload is correct, Err in db"})
				}

			} else {
				c.JSON(500, gin.H{"status": "Unable to create Enquiry"})
			}
		}
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/enquire/list", func(c *gin.Context) {
		data, _ := database.ListKeys("")
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	r.GET("/enquire/:id", func(c *gin.Context) {
		userId := c.Param("id")
		data, _ := database.Get(userId)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	r.DELETE("/enquire/:id", func(c *gin.Context) {
		userId := c.Param("id")
		database.Delete(userId)
		c.JSON(http.StatusOK, gin.H{
			"message": "Deleted data for: " + userId,
		})
	})

	// ------------
	r.POST("/users", func(c *gin.Context) {
		dt, _ := ioutil.ReadAll(c.Request.Body)

		data, errFB := services.AddUser(string(dt))
		if errFB != nil {
			c.JSON(500, gin.H{"status": errFB})
			return
		}
		c.JSON(200, data)
	})

	r.GET("/users/all", func(c *gin.Context) {
		data, errFB := services.GetAllUsers()
		if errFB != nil {
			c.JSON(500, gin.H{"status": errFB})
			return
		}
		c.JSON(200, data)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		data, errFB := services.GetOneUser(id)
		if errFB != nil {
			c.JSON(500, gin.H{"status": errFB})
			return
		}
		c.JSON(200, data)
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		data, errFB := services.DeleteOneUser(id)
		if errFB != nil {
			c.JSON(500, gin.H{"status": errFB})
			return
		}
		c.JSON(200, data)
	})
	// ------------

	r.Run()
}

//
// 700 > processor
// 8GB >
// 5000 >
