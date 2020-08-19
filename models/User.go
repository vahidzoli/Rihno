package models

import (
	"Rihno/config"
	"Rihno/entity"
	"Rihno/request"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]entity.User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *entity.User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *entity.User, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateUser(user *entity.User, req request.UserUpdate) (err error) {
	user.Name = req.Name
	pass, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	user.Password = string(pass)
	config.DB.Model(&user).Updates(req)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *entity.User) (err error) {
	config.DB.Delete(&user)
	return nil
}
