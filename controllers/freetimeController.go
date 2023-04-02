package controllers

import (
	"fmt"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFreetime(c *gin.Context) {
	var ft models.FreetimeModel
	var receivedItem models.Freetime

	if err := c.BindJSON(&receivedItem); err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "could not receive Item"})
		return
	}

	err := ft.Insert(receivedItem)
	if err != nil {
		fmt.Println("Error", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not create item"})
		return
	}

	c.IndentedJSON(http.StatusCreated, receivedItem)
}

func ReadFreetimeByUsername(c *gin.Context) {
	var ft models.FreetimeModel
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

func UpdateFreetimeByUsername(c *gin.Context) {
	var ft models.FreetimeModel
	var receivedItem models.Freetime

	if err := c.BindJSON(&receivedItem); err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "could not receive Item"})
	}

	username := c.Param("username")

	err := ft.UpdateByUsername(username, receivedItem)
	fmt.Println(receivedItem)
	if err != nil {
		fmt.Println("Error: ", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not update item"})
		return
	}

	c.IndentedJSON(http.StatusOK, receivedItem)
}
