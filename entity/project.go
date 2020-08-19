package entity

type Project struct {
	Id           uint64 `gorm:"id"`
	Name         string `gorm:"name"`
	ApiKey       string `gorm:"apiKey"`
	CallbackUrl  string `gorm:"callbackUrl"`
	VPNTunnel    bool   `gorm:"vpnTunnel"`
	UserId       uint64 `gorm:"userId"`
	Plans []Plan        `gorm:"foreignkey:ProjectId"`
}

func (b *Project) TableName() string {
	return "projects"
}
