package controllers

import (
	"fmt"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var im models.UserModel
	var receivedItem models.User

	if err := c.BindJSON(&receivedItem); err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "could not receive Item"})
		return
	}

	err := im.Insert(receivedItem)
	if err != nil {
		fmt.Println("Error", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not create item"})
		return
	}

	c.IndentedJSON(http.StatusCreated, receivedItem)
}

func ReadUserByUsername(c *gin.Context) {
	var im models.UserModel
	var err error

	username := c.Param("username")
	if username == "" {
		fmt.Println("Error: couldn't get username\n", err)
		return
	}

	receivedItem, err := im.SelectByUsername(username)
	if err != nil {
		fmt.Println("Error: ", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not find object"})
		return
	}

	c.IndentedJSON(http.StatusOK, receivedItem)
}

func ReadUsers(c *gin.Context) {
	var im models.UserModel

	items, err := im.SelectAll()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not find items"})
	}
	c.IndentedJSON(http.StatusOK, items)
}

func UpdateUserByUsername(c *gin.Context) {
	var im models.UserModel
	var receivedItem models.User

	if err := c.BindJSON(&receivedItem); err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "could not receive Item"})
	}

	username := c.Param("username")
	

	err := im.UpdateByUsername(username, receivedItem)
	if err != nil {
		fmt.Println("Error: ", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not update item"})
		return
	}

	c.IndentedJSON(http.StatusOK, receivedItem)
}
