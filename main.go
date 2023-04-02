package main

import (
	"fmt"
	"log"
	"main/controllers"
	"main/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	var err error

	err = godotenv.Load("env.env")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.Use(CORSMiddleware())

	router.GET("/users", controllers.ReadUsers)
	router.GET("/users/:username", controllers.ReadUserByUsername)
	router.POST("/users", controllers.CreateUser)
	router.PUT("/users/:username", controllers.UpdateUserByUsername)

	router.GET("/profiles", controllers.ReadProfiles)
	router.GET("/profiles/first/:firstname", controllers.ReadProfileByFirstname)
	router.GET("/profiles/last/:lastname", controllers.ReadProfileByLastname)
	router.GET("/profiles/:username", controllers.ReadProfileByUsername)
	router.POST("/profiles", controllers.CreateProfile)
	router.PUT("/profiles/first/:firstname", controllers.UpdateProfileByFirstname)
	router.PUT("/profiles/last/:lastname", controllers.UpdateProfileByLastname)

	router.GET("/freetime/:username", controllers.ReadFreetimeByUsername)
	router.POST("/freetime", controllers.CreateFreetime)
	router.PUT("/freetime/:username", controllers.UpdateFreetimeByUsername)

	router.GET("/meetings/:username", controllers.ReadMeetingByUsername)
	router.POST("/meetings", controllers.CreateMeeting)

	db.Init()

	err = router.Run("localhost:5050")
	if err != nil {
		fmt.Println("Error in launching router", err)
	}
}
