package models

import (
	"time"
	"video_application/server/database"
)

// 视频
/*
CREATE TABLE `video` (
`v_id` INT AUTO_INCREMENT PRIMARY KEY,
`u_id` INT NOT NULL,
`title` VARCHAR(50) NOT NULL,
`introduce` VARCHAR(300) COMMENT '介绍',
`video_url` VARCHAR(150) NOT NULL,
`cover` VARCHAR(200) COMMENT '封面',
`type` INT NOT NULL COMMENT '分类',
`tags` VARCHAR(30),
`status` INT DEFAULT 0 COMMENT '状态 默认为0,未通过1',
`r_id` INT COMMENT '审核人员id',
`create_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
`delete_at` DATETIME

) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8;
*/
// todo 外键 uid,rid

// 分类
const (
	ShengHuo  = 100+iota
	KeJi
	GaoXiao
	WuDao
	ShiShang
	ZiXun
	YouXi
	QiChe
	DonGman
	YingShi
	YinYue
	QiTa
)


type Video struct {
	Vid int64			`gorm:"column:v_id" json:"v_id"`
	Uid int64			`gorm:"column:u_id" json:"u_id"`
	Title string		`gorm:"column:title" json:"title"`
	Introduce string	`gorm:"column:introduce" json:"introduce,omitempty"`
	VideoUrl string		`gorm:"column:video_url" json:"video_url"`
	Cover string		`gorm:"column:cover" json:"cover,omitempty"`
	Type int			`gorm:"column:type" json:"type"`
	Tags []string		`gorm:"column:tags" json:"tags,omitempty"`

	VideoStatus int		`gorm:"column:status" json:"status,omitempty"`
	Rid	 int64			`gorm:"column:r_id" json:"-"`

	CreateAt  time.Time `gorm:"column:create_at" json:"create_at,omitempty"`
	DeleteAt  time.Time `gorm:"column:delete_id" json:"delete_id,omitempty"`
}

func (v Video) TableName() string {
	return "video"
}



// 数据库 视频数据插入
func (v *Video)NewVideo(uid int64)error{
	SQL := "INSERT INTO video (u_id,title,introduce,video_url,cover,type,tags) VALUES (?,?,?,?,?,?,?)"
	return database.DB.Exec(SQL,uid,v.Title,v.Introduce,v.VideoUrl,v.Cover,v.Type,v.Tags).Error
}

// 删除视频 添加删除时间
func (v *Video)DelVideo() error {
	SQL := "UPDATE video SET delete_id = now() WHERE v_id = ? ,u_id = ?"
	return database.DB.Exec(SQL,v.Vid,v.Uid).Error
}

// 返回类似video
func ReturnSimliarVideo(vid int64) (vides []Video){
	video := ResultVideo(vid)

	SQL := "SELECT * FROM video WHERE v_id = ? AND type = ? AND tags IN (?) ORDER BY DESC LIMIT 20"
	database.DB.Raw(SQL,vid,video.Type,video.Tags).Scan(&vides)
	return
}










type videoScan struct {
	video Video
	Number int
}

type videos struct{
	Len int
	videos []videoScan
}

// 搜索
func GetVideo(key string,page,number int) (vs videos,err error){

	SQL := "SELECT v.*,count(*) as number FROM video AS v " +
		"INNER JOIN likes AS l" +
		"v.v_id = l.v_id " +
		"WHERE v.title LIKE '%?%' AND v.type LIKE '%?%' AND v.delete_id IS NULL ORDER BY DESC LIMIT ?,?"

	err = database.DB.Raw(SQL,key,key,page,number).Scan(&vs).Error
	vs.Len = len(vs.videos)
	return
}









// 返回视频对象
/*
	或者每次在视频上传时传到缓存中,key:vid value:序列化视频对象,有改动则进行缓存删除,重置
 */
func ResultVideo(vid int64) (v Video) {
	SQL := "SELECT * FROM video WHERE v_id = ?"
	database.DB.Raw(SQL,vid).Scan(&v)
	return
}