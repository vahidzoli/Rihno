package models

import (
	"Rihno/config"
	"Rihno/entity"
	"Rihno/request"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//GetAllProjects Fetch all project data
func GetAllProjects(project *[]entity.Project, id uint64) (err error) {
	if err = config.DB.Where("user_id = ?", id).Find(project).Error; err != nil {
		return err
	}
	return nil
}

//CreateProject ... Insert New data
func CreateProject(project *entity.Project) (err error) {
	if err = config.DB.Create(project).Error; err != nil {
		return err
	}
	return nil
}

//GetProjectByID ... Fetch only one project by Id
func GetProjectByID(project *entity.Project, id string, ctx *gin.Context) (err error) {
	q := config.DB.Where("id = ?", id)
	q = userScope(q, ctx)
	if err = q.First(project).Error; err != nil {
		return err
	}
	return nil
}

//UpdateProject ... Update project
func UpdateProject(project *entity.Project, req request.Project) (err error) {
	//project.Name = req.Name
	if err = config.DB.Model(&project).Updates(req).Error; err != nil {
		return err
	}
	return nil
}

//DeleteProject ... Delete project
func DeleteProject(project *entity.Project) (err error) {
	config.DB.Delete(&project)
	return nil
}

func userScope(db *gorm.DB, ctx *gin.Context) *gorm.DB {
	if ctx != nil {
	user, _ := ctx.Get("authData")
		var userId = user.(*config.Claims).UserId
		return db.Where("user_id = ?", userId)
	}
	return nil
}