package request

type LookVideo struct {
	Vid int64			`json:"v_id" form:"v_id" binding:"required"`
	LikeVideo int 		`json:"like" form:"like"`
}