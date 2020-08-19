package controllers

import (
	"Rihno/config"
	"Rihno/entity"
	"Rihno/models"
	"Rihno/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	println("user claims:", c.Value("props"))
	var user []entity.User
	err := models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user entity.User
	err := models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, "User Not Exist!")
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var userUpdateRequest request.UserUpdate
	if err := c.ShouldBind(&userUpdateRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	var user entity.User
	id := c.Params.ByName("id")
	err := models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = models.UpdateUser(&user, userUpdateRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Something Went Wrong... :(")
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user entity.User
	id := c.Param("id")
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	err := models.DeleteUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"User " + id: "is deleted"})
	}
}
