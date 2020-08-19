package controllers

import (
	"Rihno/config"
	"Rihno/entity"
	"Rihno/models"
	"Rihno/request"
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

//GetProjects ... Get all projects
func GetProjects(c *gin.Context) {
	user, _ := c.Get("authData")
	var userId = user.(*config.Claims).UserId
	var project []entity.Project
	err := models.GetAllProjects(&project, userId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, project)
	}
}

//CreateProject ... Create project
func CreateProject(c *gin.Context) {
	var project entity.Project
	a := make([]byte, 4)
	_, err := rand.Read(a)
	if err != nil {
		panic(err)
	}
	project.Name = os.Getenv("APP_NAME")+"_"+fmt.Sprintf("%X", a)
	b := make([]byte, 8)
	_, err2 := rand.Read(b)
	if err2 != nil {
		log.Fatal(err2)
	}
	project.ApiKey = fmt.Sprintf("%X", b)
	user, _ := c.Get("authData")
	project.UserId = user.(*config.Claims).UserId

	err3 := models.CreateProject(&project)
	if err3 != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, "Project Created Successfully")
	}
}

//GetProjectByID ... Get the project by id
func GetProjectByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var project entity.Project
	err := models.GetProjectByID(&project, id, c)
	if err != nil {
		c.JSON(http.StatusNotFound, "Project Not Found!")
	} else {
		c.JSON(http.StatusOK, project)
	}
}

//UpdateProject ... Update the project information
func UpdateProject(c *gin.Context) {
	var projectUpdateRequest request.Project
	if err := c.ShouldBind(&projectUpdateRequest); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}
	var project entity.Project
	id := c.Params.ByName("id")
	err2 := models.GetProjectByID(&project, id, c)
	if err2 != nil {
		c.JSON(http.StatusNotFound, "Project Not Found!")
		return
	}
	c.BindJSON(&project)
	err3 := models.UpdateProject(&project, projectUpdateRequest)
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, "Something Went Wrong... :(")
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Project Updated Successfully"})
	}
}

//DeleteProject ... Delete the project
func DeleteProject(c *gin.Context) {
	var project entity.Project
	id := c.Param("id")
	err := models.GetProjectByID(&project, id, c)
	if err != nil {
		c.JSON(http.StatusNotFound, "Project Not Found!")
		return
	}
	err2 := models.DeleteProject(&project)
	if err2 != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"Project " + id: "is deleted"})
	}
}
