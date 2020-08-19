package request

type CreatePlan struct {
	Format    string `form:"format" json:"format" xml:"format"  binding:"required"`
	Codec 	  string `form:"codec" json:"codec" xml:"codec" binding:"required"`
}

type UpdatePlan struct {
	Name      string `form:"name" json:"name" xml:"name"`
	Format    string `form:"format" json:"format" xml:"format"`
	Codec 	  string `form:"codec" json:"codec" xml:"codec" binding:"oneof=H.264 H.265 AVI"`
}