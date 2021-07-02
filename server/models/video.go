package models



import "time"

// 视频
/*
CREATE TABLE `video` (
`vid` INT AUTO_INCREMENT PRIMARY KEY,
`uid` INT NOT NULL,
`title` VARCHAR(50) NOT NULL,
`introduce` VARCHAR(300) COMMENT '介绍',
`video_url` VARCHAR(150) NOT NULL,
`cover` VARCHAR(200) COMMENT '封面',
`type` VARCHAR(10) NOT NULL COMMENT '分类',
`tags` VARCHAR(30),
`status` INT DEFAULT 0 COMMENT '状态 默认为0,未通过1',
`rid` INT COMMENT '审核人员id',
`create_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
`delete_at` DATETIME



) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8;
*/
// todo 外键 uid,rid

type Video struct {
	Vid int64			`gorm:"column:vid" json:"vid"`
	Uid int64			`gorm:"column:uid" json:"uid"`
	Title string		`gorm:"column:title" json:"title"`
	Introduce string	`gorm:"column:introduce" json:"introduce,omitempty"`
	VideoUrl string		`gorm:"column:video_url" json:"video_url"`
	Cover string		`gorm:"column:cover" json:"cover,omitempty"`
	Type string			`gorm:"column:type" json:"type"`
	Tags []string		`gorm:"column:tags" json:"tags,omitempty"`

	VideoStatus int		`gorm:"column:status" json:"status,omitempty"`
	Rid	 int64			`gorm:"column:rid" json:"-"`

	CreateAt  time.Time `gorm:"column:create_at" json:"create_at,omitempty"`
	DeleteAt  time.Time `gorm:"column:delete_id" json:"delete_id,omitempty"`
}

func (b Video) TableName() string {
	return "video"
}


