package models

import (
	"Rihno/config"
	"Rihno/entity"
)

func GetAllResolutions(resolution *[]entity.Resolution) (err error) {
	if err = config.DB.Find(resolution).Error; err != nil {
		return err
	}
	return nil
}
