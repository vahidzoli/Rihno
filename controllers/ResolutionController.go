package controllers

import (
	"Rihno/entity"
	"Rihno/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetResolutions ... Get all Resolutions
func GetResolutions(c *gin.Context) {
	var resolution []entity.Resolution
	err := models.GetAllResolutions(&resolution)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"Data": resolution})
	}
}
