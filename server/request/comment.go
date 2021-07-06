package request

type Comment struct {
	Cid int64			`gorm:"column:c_id" json:"c_id"`	// 评论主键id
	Vid int64			`gorm:"column:v_id" json:"v_id"`	// 视频id
	// 回复类型 视频评论0 还是回复他人的评论1
	ComposeType int		`gorm:"column:compose_type" json:"compose_type"`
	// 评论内容
	Content	 string		`gorm:"column:content" json:"content"`
	// 评论用户id
	FromUid int64 		`gorm:"column:from_uid" json:"from_uid"`
	// 目标用户id
	ToUid int64 		`gorm:"column:to_uid" json:"to_uid"`
}
