package models



import "time"

// 观看视频或进行点赞
/*
CREATE TABLE `view_video` (
`vid` INT NOT NULL,
`uid` INT NOT NULL,
`like` INT DEFAULT 0 COMMENT '点赞行为,默认不点赞为0,点赞为1',
`create_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
`delete_at` DATETIME
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
// todo 外键

type ViewVideo struct {
	Vid int64			`gorm:"column:vid" json:"vid"`
	Uid int64			`gorm:"column:uid" json:"uid"`
	LikeVideo int 		`gorm:"column:like" json:"like"`
	CreateAt  time.Time `gorm:"column:create_at" json:"create_at,omitempty"`
	DeleteAt  time.Time `gorm:"column:delete_id" json:"delete_id,omitempty"`
}

func (b ViewVideo) TableName() string {
	return "view_video"
}