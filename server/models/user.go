package models


import "time"

// 用户
/*
CREATE TABLE `user` (
`uid` INT AUTO_INCREMENT PRIMARY KEY,
`username` VARCHAR(20) NOT NULL,
`password` VARCHAR(150) NOT NULL,
`nickname` VARCHAR(20) NOT NULL,
`portrait` VARCHAR(200) COMMENT '头像',
`introduce` VARCHAR(300) COMMENT '用户介绍',
`create_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
`delete_at` DATETIME
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8;
*/

type User struct {
	Uid int64			`gorm:"column:uid" json:"uid"`
	UserName string		`gorm:"column:username" json:"username"`
	PassWord string		`gorm:"column:password" json:"password"`
	NickName string		`gorm:"column:nickname" json:"nickname"`
	Portrait string		`gorm:"column:portrait" json:"portrait,omitempty"`
	Introduce string	`gorm:"column:introduce" json:"introduce,omitempty"`

	CreateAt  time.Time `gorm:"column:create_at" json:"create_at,omitempty"`
	DeleteAt  time.Time `gorm:"column:delete_id" json:"delete_id,omitempty"`
}

func (b User) TableName() string {
	return "user"
}

