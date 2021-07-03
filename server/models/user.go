package models

import (
	"time"
	"video_application/server/database"
)

// 用户
/*
CREATE TABLE `user` (
`u_id` INT AUTO_INCREMENT PRIMARY KEY,
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
	Uid int64			`gorm:"column:u_id" json:"u_id"`
	UserName string		`gorm:"column:user_name" json:"user_name"`
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




// 查询密码
func (u *User)GetUserPassword()  {
	database.DB.Where("user_name = ? ",u.UserName).First(u)
}

// 查询用户账号
func (u *User)GetUserName()  {
	database.DB.Where("user_name = ? ",u.UserName).First(u)
}

// 根据用户id查询
func (u *User)GetToUidUser(){
	database.DB.Where("u_id = ? ",u.Uid).First(u)
}

// 注册
func (u *User)Registry() error {
	sql := "INSERT INTO user (nickname,user_name,password,) VALUES (?,?,?)"
	return database.DB.Raw(sql,u.NickName,u.UserName,u.PassWord).Error
}


// 修改用户资料
func (u *User)ModifyUserData() error {
	SQL := "UPDATE user SET nickname = ?,portrait=?,introduce=? WHERE u_id = ?"
	return database.DB.Exec(SQL,u.NickName,u.Portrait,u.Introduce,u.Uid).Error
}

// 修改用户密码
func (u *User)ModifyUserPassword() error {
	SQL := "UPDATE user SET password = ? WHERE u_id = ?"
	return database.DB.Exec(SQL,u.PassWord,u.Uid).Error
}






