package models

import (
	"errors"
	"time"
	"video_application/server/database"
)

// 评论
/*
CREATE TABLE `comments` (
`c_id` INT AUTO_INCREMENT PRIMARY KEY,
`v_id` INT NOT NULL,
//   `compose_type` INT NOT NULL COMMENT '回复类型 视频评论0 还是回复他人的评论1',
`content` VARCHAR(200) NOT NULL,
`from_uid` INT NOT NULL COMMENT '评论用户id',
//`likes` INT DEFAULT 0 COMMENT '点赞数',
`create_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
`delete_at` DATETIME

) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8;
*/

// todo 点赞数对数据库太过于频繁,可以点赞时写入缓存中,后台运行每隔一段时间获取点赞总数并删除缓存数据写入库中
type Comments struct {
	Cid int64			`gorm:"column:c_id" json:"c_id"`	// 评论id 主键
	Vid int64			`gorm:"column:v_id" json:"v_id"`	// 视频id

	// 回复类型 视频评论0 还是回复他人的评论1
	//ComposeType int		`gorm:"column:compose_type" json:"compose_type"`

	// 评论内容
	Content	 string		`gorm:"column:content" json:"content"`

	// 评论用户id
	FromUid int64 		`gorm:"column:from_uid" json:"from_uid"`
	//Likes int64			`gorm:"column:likes" json:"likes"`	// 点赞数

	CreateAt  time.Time `gorm:"column:create_at" json:"create_at,omitempty"`
	DeleteAt  time.Time `gorm:"column:delete_id" json:"delete_id,omitempty"`
}

func (b Comments) TableName() string {
	return "comments"
}

// 添加评论
func (b *Comments) AddComment()  error {
	SQL := "INSERT INTO comments (v_id,content,from_uid) VALUES (?,?,?)"
	return database.DB.Exec(SQL,b.Vid,b.Content,b.FromUid).Error
}



/*
CREATE TABLE `reply_comment` (
`id` INT AUTO_INCREMENT PRIMARY KEY,
`c_id` INT NOT NULL COMMENT '评论id',
//`reply_id` INT NOT NULL COMMENT '回复目标id',
//`ReplyType` INT NOT NULL COMMENT '回复类型',
`content` VARCHAR(200) NOT NULL COMMENT '回复内容',
`from_uid` INT NOT NULL COMMENT '回复用户id',
`to_uid` INT NOT NULL COMMENT '目标用户id',
//`likes` INT DEFAULT 0 COMMENT '点赞数',
`create_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
`delete_at` DATETIME
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8;
*/
// 回评表
type ReplyComment struct {
	Id	int64 			`gorm:"column:id" json:"id"`	//主键
	Cid	int64 			`gorm:"column:c_id" json:"c_id"`	//评论id
	//ReplyId	int64 		`gorm:"column:reply_id" json:"reply_id"`	//回复目标id
	//ReplyType int		`gorm:"column:reply_type" json:"reply_type"`	//回复类型
	Content	string 		`gorm:"column:content" json:"content"`	//回复内容
	FromUid	int64 		`gorm:"column:from_uid" json:"from_uid"`	// 回复用户id
	ToUid int64 		`gorm:"column:to_uid" json:"to_uid"`	// 目标用户id
	//Likes int64			`gorm:"column:likes" json:"likes"`	// 点赞数
	CreateAt  time.Time `gorm:"column:create_at" json:"create_at,omitempty"`
	DeleteAt  time.Time `gorm:"column:delete_id" json:"delete_id,omitempty"`
}

func (b ReplyComment) TableName() string {
	return "reply_comment"
}

// 添加回复评论
func (b *ReplyComment) AddReplyComment() error {
	SQL := "INSERT INTO reply_comment (c_id,from_uid,to_uid,content) VALUES (?,?,?,?)"
	return database.DB.Exec(SQL,b.Cid,b.FromUid,b.ToUid,b.Content).Error
}

// 删除评论
func DelComment(uid,cid int64,isreply int)(err error){
	var SQL string
	if isreply == 0 {
		SQL = "UPDATE comments SET delete_id = now() WHERE c_id = ? AND u_id = ?"
		err = database.DB.Exec(SQL,cid,uid).Error
	}else if isreply == 1 {
		SQL = "UPDATE reply_comment SET delete_id = now() WHERE c_id = ? AND u_id = ?"
		err = database.DB.Exec(SQL,cid,uid).Error
	}

	return
}



/*
CREATE TABLE `likes` (
//`id` INT AUTO_INCREMENT PRIMARY KEY,
//`type_id` INT NOT NULL COMMENT '对应评论的id',
`u_id` INT NOT NULL,
`c_id` INT NOT NULL,
`v_id` INT NOT NULL,
`status` INT NOT NULL COMMENT '点赞状态  0--默认无赞   1--有效赞'
//) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8;
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

// 点赞
type Likes struct {
	//Id int64			`gorm:"column:id" json:"id"`
	//TypeId int64		`gorm:"column:type_id" json:"type_id"`	//对应评论的id
	Uid int64			`gorm:"column:u_id" json:"u_id"`
	Vid int64			`gorm:"column:v_id" json:"v_id"`	// 视频id
	Cid	int64 			`gorm:"column:c_id" json:"c_id"`	//评论id

	Status int			`gorm:"column:status" json:"status"`	// 点赞状态  0--默认无赞   1--有效赞
}
func (b Likes) TableName() string {
	return "likes"
}

// 增加点赞
func (b *Likes) AddLike(uid int64) (err error) {
	SQL := "INSERT INTO like_comment (u_id,v_id,c_id,status) VALUES (?,?,?,?)"
	return database.DB.Exec(SQL,uid,b.Vid,b.Cid,b.Status).Error
}

// 删除点赞
func (b *Likes) DelLike(uid int64)  (err error) {
	var SQL string

	if b.Vid != 0 && b.Cid == 0 {
		SQL = "UPDATE likes SET v_id = ? ,status = ? WHERE u_id = ?"
		err = database.DB.Exec(SQL,b.Vid,b.Status,uid).Error
	}else if b.Vid == 0 && b.Cid != 0 {
		SQL = "UPDATE likes SET c_id = ? ,status = ? WHERE u_id = ?"
		err = database.DB.Exec(SQL,b.Cid,b.Status,uid).Error
	}
	return 	errors.New("点赞异常")
}