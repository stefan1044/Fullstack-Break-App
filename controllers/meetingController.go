package controllers

import (
	"fmt"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMeeting(c *gin.Context) {
	var mt models.MeetingModel
	var receivedItem models.Meeting

	if err := c.BindJSON(&receivedItem); err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "could not receive Item"})
		return
	}

	err := mt.Insert(receivedItem)
	if err != nil {
		fmt.Println("Error", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not create item"})
		return
	}

	c.IndentedJSON(http.StatusCreated, receivedItem)
}

func ReadMeetingByUsername(c *gin.Context) {
	var ft models.MeetingModel
	var err error

	username := c.Param("username")
	if username == "" {
		fmt.Println("Error: couldn't get username\n", err)
		return
	}

	receivedItem, err := ft.SelectByUsername(username)
	if err != nil {
		fmt.Println("Error: ", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not find object"})
		return
	}

	c.IndentedJSON(http.StatusOK, receivedItem)
}
