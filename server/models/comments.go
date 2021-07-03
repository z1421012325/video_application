package models


import "time"

// 评论
/*
CREATE TABLE `comments` (
`c_id` INT AUTO_INCREMENT PRIMARY KEY,
`v_id` INT NOT NULL,
`compose_type` INT NOT NULL COMMENT '回复类型 视频评论0 还是回复他人的评论1',
`content` VARCHAR(200) NOT NULL,
`from_uid` INT NOT NULL COMMENT '评论用户id',
`create_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
`delete_at` DATETIME

) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8;
*/

type Comments struct {

	Cid int64			`gorm:"column:c_id" json:"c_id"`	// 评论主键id
	Vid int64			`gorm:"column:v_id" json:"v_id"`	// 视频id
	// 回复类型 视频评论0 还是回复他人的评论1
	ComposeType int		`gorm:"column:compose_type" json:"compose_type"`
	// 评论内容
	Content	 string		`gorm:"column:content" json:"content"`
	// 评论用户id
	FromUid int64 		`gorm:"column:from_uid" json:"from_uid"`
	CreateAt  time.Time `gorm:"column:create_at" json:"create_at,omitempty"`
	DeleteAt  time.Time `gorm:"column:delete_id" json:"delete_id,omitempty"`
}

func (b Comments) TableName() string {
	return "comments"
}







/*
CREATE TABLE `reply_comment` (
`id` INT AUTO_INCREMENT PRIMARY KEY,
`c_id` INT NOT NULL COMMENT '评论id',
`reply_id` INT NOT NULL COMMENT '回复目标id',
`ReplyType` INT NOT NULL COMMENT '回复类型',
`content` VARCHAR(200) NOT NULL COMMENT '回复内容',
`from_uid` INT NOT NULL COMMENT '回复用户id',
`to_uid` INT NOT NULL COMMENT '目标用户id',
`create_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
`delete_at` DATETIME
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8;
*/
// 回评表
type ReplyComment struct {
	Id	int64 			`gorm:"column:id" json:"id"`	//主键
	Cid	int64 			`gorm:"column:c_id" json:"c_id"`	//评论id
	ReplyId	int64 		`gorm:"column:reply_id" json:"reply_id"`	//回复目标id
	ReplyType int		`gorm:"column:reply_type" json:"reply_type"`	//回复类型
	Content	string 		`gorm:"column:content" json:"content"`	//回复内容
	FromUid	int64 		`gorm:"column:from_uid" json:"from_uid"`	// 回复用户id
	ToUid int64 		`gorm:"column:to_uid" json:"to_uid"`	// 目标用户id
	CreateAt  time.Time `gorm:"column:create_at" json:"create_at,omitempty"`
	DeleteAt  time.Time `gorm:"column:delete_id" json:"delete_id,omitempty"`
}

func (b ReplyComment) TableName() string {
	return "reply_comment"
}








/*
CREATE TABLE `like_comment` (
`id` INT AUTO_INCREMENT PRIMARY KEY,
`type_id` INT NOT NULL COMMENT '对应评论的id',
`u_id` INT NOT NULL,
`status` INT NOT NULL COMMENT '点赞状态  0--默认无赞   1--有效赞'
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8;
*/
// 评论点赞
type LikeComment struct {
	Id int64			`gorm:"column:id" json:"id"`
	TypeId int64		`gorm:"column:type_id" json:"type_id"`	//对应评论的id
	Uid int64			`gorm:"column:u_id" json:"u_id"`
	Status int			`gorm:"column:status" json:"status"`	// 点赞状态  0--默认无赞   1--有效赞

}
func (b LikeComment) TableName() string {
	return "like_comment"
}