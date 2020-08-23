package entity

type Plan struct {
	Id          uint64        `gorm:"id"`
	Name        string        `gorm:"name"`
	UniqueKey   string        `gorm:"uniqueKey"`
	Format      string        `gorm:"format"`
	Codec       string        `gorm:"codec"`
	ProjectId   uint64        `gorm:"projectId"`
	Resolutions []*Resolution `gorm:"many2many:plan_resolutions;association_autoupdate:false"`
}

func (b *Plan) TableName() string {
	return "plans"
}
