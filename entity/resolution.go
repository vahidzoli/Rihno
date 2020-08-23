package entity

type Resolution struct {
	Id     uint64  `gorm:"id"`
	Name   string  `gorm:"name"`
	Width  string  `gorm:"width"`
	Height string  `gorm:"height"`
	Plans  []*Plan `gorm:"many2many:plan_resolutions;" json:"-"`
}

func (b *Resolution) TableName() string {
	return "resolutions"
}
