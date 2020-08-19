package entity

type User struct {
	Id      	uint64 `gorm:"id"`
	Name    	string `gorm:"name"`
	Email   	string `gorm:"email"`
	Password   	string `gorm:"password"`
	Projects []Project `gorm:"foreignkey:UserId"`
}

func (b *User) TableName() string {
	return "users"
}
