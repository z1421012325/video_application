package request



// 增加评论结构体
type Comment struct {
	Cid int64			`json:"c_id" form:"c_id"`	// 评论id
	Vid int64			`json:"v_id" form:"v_id"`	// 视频id
	// 回复类型 视频评论0 还是回复他人的评论1       gt,lt 限制参数在0~1
	ComposeType int		`json:"compose_type" form:"compose_type" binding:"gt=0,lt=1"`
	// 评论内容
	Content	 string		`json:"content" form:"content"`
	// 评论用户id
	FromUid int64 		`json:"from_uid" form:"from_uid"`
	// 目标用户id
	ToUid int64 		`json:"to_uid" form:"to_uid"`
}

// 删除评论
type DelComment struct {
	Cid int64			`json:"c_id" form:"c_id" binding:"required"`	// 评论id
	// 是否确定视频回复 0   回评回复 1
	IsReply int			`json:"is_reply" form:"is_reply" binding:"required,gt=0,lt=1"`
}



// 点赞结构体
type Like struct {
	Cid int64			`json:"c_id" form:"c_id"`	// 评论id
	Vid int64			`json:"v_id" form:"v_id"`	// 视频id
	// 回复类型 视频评论0 还是回复他人的评论1       gt,lt 限制参数在0~1
	//ComposeType int		`json:"compose_type" form:"compose_type" binding:"gt=0,lt=1"`
	// 点赞or取消点赞      取消点赞 0             点赞 1
	Status int			`json:"status" form:"status" binding:"gt=0,lt=1"`
}