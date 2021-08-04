package request




// 视频点击
type LookVideo struct {
	Vid int64			`json:"v_id" form:"v_id" binding:"required"`
	LikeVideo int 		`json:"like" form:"like"`
}


// 搜索
type Search struct {
	Key string			`json:"key" form:"key"`
	Page int			`json:"page" form:"key" binding:"gt=0"`
	Number int			`json:"number" form:"number" binging:"gt=1,lt=100"`
}
