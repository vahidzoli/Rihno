package models

import (
	"Rihno/config"
	"Rihno/entity"
	"Rihno/request"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllPlans(Plan *[]entity.Plan, project entity.Project) (err error) {
	if err = config.DB.Where("project_id = ?",project.Id).Find(Plan).Error; err != nil {
		return err
	}
	return nil
}

//CreatePlan ... Insert New data
func CreatePlan(Plan *entity.Plan) (err error) {
	if err = config.DB.Create(Plan).Error; err != nil {
		return err
	}
	return nil
}

//GetPlanByID ... Fetch only one Plan by Id
func GetPlanByID(Plan *entity.Plan, id string, project entity.Project) (err error) {
	if err = config.DB.Where("id = ?", id).Where("project_id = ?", project.Id).First(Plan).Error; err != nil {
		return err
	}
	return nil
}

//UpdatePlan ... Update Plan
func UpdatePlan(Plan *entity.Plan, req request.UpdatePlan) (err error) {
	config.DB.Model(&Plan).Updates(req)
	return nil
}

//DeletePlan ... Delete Plan
func DeletePlan(Plan *entity.Plan) (err error) {
	config.DB.Delete(&Plan)
	return nil
}
