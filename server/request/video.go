package request

import "video_application/server/models"


type UpVideo struct {
	Uid int64			`json:"u_id" form:"u_id" binding:"required"`
	Title string		`json:"title" form:"title" binding:"required,max=50"`
	Introduce string	`json:"introduce" form:"introduce" binding:"max=300"`
	VideoUrl string		`json:"video_url" form:"video_url" binding:"required,min=2,max=150"`
	Cover string		`json:"cover" form:"cover" binding:"min=2,max=200"`
	Type int			`json:"type" form:"type"`
	Tags []string		`json:"tags" form:"tags" binding:"max=30"`
}

// 类型检查
func (v *UpVideo)CheckVideoType() {
	switch v.Type {
	case models.ShengHuo :
	case models.KeJi :
	case models.GaoXiao :
	case models.WuDao :
	case models.ShiShang :
	case models.ZiXun :
	case models.YouXi :
	case models.QiChe :
	case models.DonGman :
	case models.YingShi :
	case models.YinYue :
	default:
		v.Type = models.QiTa
	}
}


type DelVideo struct {
	Uid int64			`json:"u_id" form:"u_id" binding:"required"`
	Vid int64			`json:"v_id" form:"v_id" binding:"required"`
}