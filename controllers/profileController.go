package controllers

import (
	"fmt"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProfile(c *gin.Context) {
	var im models.PersonModel
	var receivedItem models.Person

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

func ReadProfiles(c *gin.Context) {
	var im models.PersonModel

	items, err := im.SelectAll()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not find items"})
	}
	c.IndentedJSON(http.StatusOK, items)
}
func ReadProfileByFirstname(c *gin.Context) {
	var im models.PersonModel
	var err error

	firstname := c.Param("firstname")
	if firstname == "" {
		fmt.Println("Error: couldn't get firstname\n", err)
		return
	}

	receivedItem, err := im.SelectByFirstname(firstname)
	if err != nil {
		fmt.Println("Error: ", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not find object"})
		return
	}

	c.IndentedJSON(http.StatusOK, receivedItem)
}
func ReadProfileByLastname(c *gin.Context) {
	var im models.PersonModel
	var err error

	firstname := c.Param("lastname")
	if firstname == "" {
		fmt.Println("Error: couldn't get firstname\n", err)
		return
	}

	receivedItem, err := im.SelectByFirstname(firstname)
	if err != nil {
		fmt.Println("Error: ", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not find object"})
		return
	}

	c.IndentedJSON(http.StatusOK, receivedItem)
}
func ReadProfileByUsername(c *gin.Context) {
	var im models.PersonModel
	var err error

	username := c.Param("username")
	if username == "" {
		fmt.Println("Error: couldn't get firstname\n", err)
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

func UpdateProfileByFirstname(c *gin.Context) {
	var im models.PersonModel
	var receivedItem models.Person

	if err := c.BindJSON(&receivedItem); err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "could not receive Item"})
	}

	username := c.Param("firstname")

	err := im.UpdateByFirstrname(username, receivedItem)
	if err != nil {
		fmt.Println("Error: ", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not update item"})
		return
	}

	c.IndentedJSON(http.StatusOK, receivedItem)
}
func UpdateProfileByLastname(c *gin.Context) {
	var im models.PersonModel
	var receivedItem models.Person

	if err := c.BindJSON(&receivedItem); err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "could not receive Item"})
	}

	username := c.Param("lastname")

	err := im.UpdateByFirstrname(username, receivedItem)
	if err != nil {
		fmt.Println("Error: ", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not update item"})
		return
	}

	c.IndentedJSON(http.StatusOK, receivedItem)
}
