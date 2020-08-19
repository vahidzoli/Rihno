package request

type UserUpdate struct {
	Name     string `form:"name" json:"name" xml:"name"`
	Password string `form:"password" json:"password" xml:"password"`
}