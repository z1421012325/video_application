package models



import "time"

// 视频
/*
CREATE TABLE `video` (
`v_id` INT AUTO_INCREMENT PRIMARY KEY,
`u_id` INT NOT NULL,
`title` VARCHAR(50) NOT NULL,
`introduce` VARCHAR(300) COMMENT '介绍',
`video_url` VARCHAR(150) NOT NULL,
`cover` VARCHAR(200) COMMENT '封面',
`type` VARCHAR(10) NOT NULL COMMENT '分类',
`tags` VARCHAR(30),
`status` INT DEFAULT 0 COMMENT '状态 默认为0,未通过1',
`r_id` INT COMMENT '审核人员id',
`create_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
`delete_at` DATETIME



) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8;
*/
// todo 外键 uid,rid

type Video struct {
	Vid int64			`gorm:"column:v_id" json:"v_id"`
	Uid int64			`gorm:"column:u_id" json:"u_id"`
	Title string		`gorm:"column:title" json:"title"`
	Introduce string	`gorm:"column:introduce" json:"introduce,omitempty"`
	VideoUrl string		`gorm:"column:video_url" json:"video_url"`
	Cover string		`gorm:"column:cover" json:"cover,omitempty"`
	Type string			`gorm:"column:type" json:"type"`
	Tags []string		`gorm:"column:tags" json:"tags,omitempty"`

	VideoStatus int		`gorm:"column:status" json:"status,omitempty"`
	Rid	 int64			`gorm:"column:r_id" json:"-"`

	CreateAt  time.Time `gorm:"column:create_at" json:"create_at,omitempty"`
	DeleteAt  time.Time `gorm:"column:delete_id" json:"delete_id,omitempty"`
}

func (b Video) TableName() string {
	return "video"
}


