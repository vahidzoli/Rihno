package controllers

import (
	"Rihno/entity"
	"Rihno/models"
	"Rihno/request"
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//GetPlans ... Get all project's plans
func GetPlans(c *gin.Context) {
	id := c.Params.ByName("id")
	var project entity.Project
	err := models.GetProjectByID(&project, id, c)
	if err != nil {
		c.JSON(http.StatusNotFound, "Project Not Found!")
		return
	}
	var plan []entity.Plan
	err2 := models.GetAllPlans(&plan, project)
	if err2 != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"plan": plan})
	}
}

//CreatePlan ... Create Plan
func CreatePlan(c *gin.Context) {
	var planRequest request.CreatePlan
	if err := c.ShouldBind(&planRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	var Plan entity.Plan
	Plan.Name = "Plan name"
	b := make([]byte, 8)
	_, err2 := rand.Read(b)
	if err2 != nil {
		log.Fatal(err2)
	}
	Plan.UniqueKey = fmt.Sprintf("%X", b)
	Plan.Format = planRequest.Format
	Plan.Codec = planRequest.Codec
	id := c.Params.ByName("id")
	var project entity.Project
	err := models.GetProjectByID(&project, id, c)
	if err != nil {
		c.JSON(http.StatusNotFound, "Project Not Found!")
		return
	}
	u64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	Plan.ProjectId = u64
	err3 := models.CreatePlan(&Plan)
	if err3 != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, "Plan Created Successfully")
	}
}

//GetPlanByID ... Get the Plan by id
func GetPlanByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var project entity.Project
	err := models.GetProjectByID(&project, id, c)
	if err != nil {
		c.JSON(http.StatusNotFound, "Project Not Found!")
		return
	}
	pid := c.Params.ByName("pid")
	var Plan entity.Plan
	err2 := models.GetPlanByID(&Plan, pid, project)
	if err2 != nil {
		c.JSON(http.StatusNotFound, "Plan Not Found!")
	} else {
		c.JSON(http.StatusOK, Plan)
	}
}

//UpdatePlan ... Update the Plan information
func UpdatePlan(c *gin.Context) {
	var planUpdateRequest request.UpdatePlan
	if err := c.ShouldBind(&planUpdateRequest); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}
	var project entity.Project
	id := c.Params.ByName("id")
	err := models.GetProjectByID(&project, id, c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Project Not Found!"})
		return
	}
	var plan entity.Plan
	pid := c.Params.ByName("pid")
	err2 := models.GetPlanByID(&plan, pid, project)
	if err2 != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Plan Not Found!"})
		return
	}
	c.BindJSON(&plan)
	err3 := models.UpdatePlan(&plan, planUpdateRequest)
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, "Something Went Wrong... :(")
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Plan Updated Successfully"})
	}
}

//DeletePlan ... Delete the Plan
func DeletePlan(c *gin.Context) {
	var project entity.Project
	id := c.Param("id")
	err := models.GetProjectByID(&project, id, c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Project Not Found!"})
		return
	}
	var plan entity.Plan
	pid := c.Param("pid")
	err2 := models.GetPlanByID(&plan, pid, project)
	if err2 != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Plan Not Found!"})
		return
	}
	err3 := models.DeletePlan(&plan)
	if err3 != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"Plan " + pid: "is deleted"})
	}
}
